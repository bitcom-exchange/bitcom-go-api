package wsclient

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/ws"
	"github.com/gorilla/websocket"
)

const (
	TimerIntervalSecond = 5

	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	// TODO: Set pong wait to a long time to avoid network problems
	pongWait = 10 * time.Minute

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = 54 * time.Second // (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 2 * 1024 * 1024

	path = "/"
)

type ConnectedHandler func()

type MessageHandler func(message string) (interface{}, error)

type ResponseHandler func(response interface{})

// When called, returns a new websocket token
type TokenProducer func() string

type WebSocketClientBase struct {
	host                string
	token               string
	tokenProducer       TokenProducer
	conn                *websocket.Conn
	connectedHandler    ConnectedHandler
	responseHandler     ResponseHandler
	stopReadChannel     chan struct{}
	ticker              *time.Ticker
	lastReceivedTime    time.Time
	reconnectWaitSecond int64

	// for WriteMessage
	sendMutex *sync.Mutex
}

func (wsc *WebSocketClientBase) Init(host string, tokenProducer TokenProducer, reconnectWaitSecond int64) *WebSocketClientBase {
	wsc.host = host
	wsc.tokenProducer = tokenProducer
	wsc.stopReadChannel = make(chan struct{})
	wsc.reconnectWaitSecond = reconnectWaitSecond
	wsc.sendMutex = &sync.Mutex{}
	return wsc
}

func (wsc *WebSocketClientBase) BuildRequestBody(paramMap map[string]interface{}) (string, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	if wsc.token != "" {
		paramMap["token"] = wsc.token
	}

	var jsonBody string

	if len(paramMap) > 0 {
		jsonBuf, err := json.Marshal(paramMap)
		if err != nil {
			return "", err
		}

		jsonBody = string(jsonBuf)
	}

	return jsonBody, nil
}

func (wsc *WebSocketClientBase) SetHandler(connHandler ConnectedHandler, repHandler ResponseHandler) {
	wsc.connectedHandler = connHandler
	wsc.responseHandler = repHandler
}

func (wsc *WebSocketClientBase) Connect(autoConnect bool) {
	wsc.connectWebSocket()

	if autoConnect {
		wsc.startTicker()
	}
}

func (wsc *WebSocketClientBase) Send(data string) {
	if wsc.conn == nil {
		applogger.Error("WebSocket sent error: no connection available")
		return
	}

	if err := wsc.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		applogger.Error("Set write dead line error")
		return
	}

	wsc.sendMutex.Lock()
	err := wsc.conn.WriteMessage(websocket.TextMessage, []byte(data))
	if err != nil {
		applogger.Error("WebSocket sent error: data=%s, error=%s", data, err)
	}
	wsc.sendMutex.Unlock()

}

func (wsc *WebSocketClientBase) Close() {
	wsc.sendMutex.Lock()
	err := wsc.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		applogger.Error("Write close:", err)
	}
	wsc.sendMutex.Unlock()

	wsc.disconnectWebSocket()
}

func (wsc *WebSocketClientBase) connectWebSocket() {
	var err error
	url := fmt.Sprintf("%s%s", wsc.host, path)
	applogger.Debug("WebSocket connecting...")
	wsc.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		applogger.Error("WebSocket connected error: %s", err)
		return
	}
	applogger.Info("WebSocket connected")

	if wsc.tokenProducer != nil {
		wsc.token = wsc.tokenProducer()
	}

	wsc.conn.SetPingHandler(nil)
	wsc.conn.SetReadLimit(maxMessageSize)

	if err := wsc.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		applogger.Error("WebSocket connected error: %s", err)
		return
	}

	wsc.conn.SetPongHandler(func(string) error {
		return wsc.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	wsc.conn.SetCloseHandler(func(code int, text string) error {
		applogger.Info("Receive close message")
		wsc.Close()
		return nil
	})

	wsc.startReadLoop()

	if wsc.connectedHandler != nil {
		wsc.connectedHandler()
	}
}

func (wsc *WebSocketClientBase) disconnectWebSocket() {
	if wsc.conn == nil {
		return
	}

	applogger.Debug("WebSocket disconnecting...")
	wsc.stopReadLoop()
	time.Sleep(time.Second)

	err := wsc.conn.Close()
	if err != nil {
		applogger.Error("WebSocket disconnect error: %s", err)
		return
	}

	applogger.Info("WebSocket disconnected")
}

func (wsc *WebSocketClientBase) startTicker() {
	wsc.ticker = time.NewTicker(TimerIntervalSecond * time.Second)
	wsc.lastReceivedTime = time.Now()

	go wsc.tickerLoop()
}

func (wsc *WebSocketClientBase) tickerLoop() {
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		wsc.ticker.Stop()
		pingTicker.Stop()
	}()

	applogger.Debug("tickerLoop started")
	for {
		select {
		// Receive tick from tickChannel
		case <-wsc.ticker.C:
			elapsedSecond := time.Now().Sub(wsc.lastReceivedTime).Seconds()
			applogger.Debug("WebSocket received data %f sec ago", elapsedSecond)

			if elapsedSecond > float64(wsc.reconnectWaitSecond) {
				applogger.Info("WebSocket reconnect...")
				wsc.ticker.Stop()
				wsc.disconnectWebSocket()
				wsc.connectWebSocket()
				wsc.ticker = time.NewTicker(TimerIntervalSecond * time.Second)

				// reset lastReceivedTime after reconnection
				wsc.lastReceivedTime = time.Now()
			}

		case <-pingTicker.C:
			if wsc.conn == nil {
				applogger.Error("Connection is null!")
				continue
			}
			_ = wsc.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := wsc.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				applogger.Error("Write ping message err :%s", err)
				wsc.disconnectWebSocket()
				wsc.connectWebSocket()
			}
		}
	}
}

func (wsc *WebSocketClientBase) startReadLoop() {
	go wsc.readLoop()
}

func (wsc *WebSocketClientBase) stopReadLoop() {
	wsc.stopReadChannel <- struct{}{}
}

func (wsc *WebSocketClientBase) readLoop() {
	applogger.Debug("readLoop started")
	for {
		select {
		case <-wsc.stopReadChannel:
			applogger.Debug("readLoop stopped")
			return

		default:
			if wsc.conn == nil {
				applogger.Error("Read error: no connection available")
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			// TODO: If no message receives, the function will be blocked here
			msgType, buf, err := wsc.conn.ReadMessage()
			if err != nil {
				applogger.Error("Read error: %s", err)
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			switch msgType {
			case websocket.TextMessage:
				wsc.lastReceivedTime = time.Now()
				message := string(buf)
				wsc.handleTextMessage(message)
			}

		}
	}
}

func (wsc *WebSocketClientBase) handleTextMessage(message string) {
	switch {
	case strings.Contains(message, "\"channel\":\"subscription\""):
		wsc.handleSubscriptionMessage(message)
	case strings.Contains(message, "\"channel\":\"depth\"") && !strings.Contains(message, "depth1"):
		wsc.handleDepthMessage(message)
	case strings.Contains(message, "\"channel\":\"order_book."):
		wsc.handleOrderBookMessage(message)
	case strings.Contains(message, "\"channel\":\"depth1\""):
		wsc.handleDepth1Message(message)
	case strings.Contains(message, "\"channel\":\"ticker\""):
		wsc.handleTickerMessage(message)
	case strings.Contains(message, "\"channel\":\"kline."):
		wsc.handleKlineMessage(message)
	case strings.Contains(message, "\"channel\":\"trade\""):
		wsc.handleTradeMessage(message)
	case strings.Contains(message, "\"channel\":\"market_trade\""):
		wsc.handleMarketTradeMessage(message)
	case strings.Contains(message, "\"channel\":\"index_price\""):
		wsc.handleIndexPriceMessage(message)
	case strings.Contains(message, "\"channel\":\"mark_price\""):
		wsc.handleMarkPriceMessage(message)
	case strings.Contains(message, "\"channel\":\"account\""):
		wsc.handleAccountMessage(message)
	case strings.Contains(message, "\"channel\":\"position\""):
		wsc.handlePositionMessage(message)
	case strings.Contains(message, "\"channel\":\"order\"") && !strings.Contains(message, "order_book"):
		wsc.handleOrderMessage(message)
	case strings.Contains(message, "\"channel\":\"user_trade\""):
		wsc.handleUserTradeMessage(message)
	}
}

func (wsc *WebSocketClientBase) handleSubscriptionMessage(message string) {
	successResponse := ws.SubscriptionSuccessResponse{}
	err := json.Unmarshal([]byte(message), &successResponse)
	if err != nil {
		applogger.Error("Handle subscription message error: %s", err)
	}

	failResponse := ws.SubscriptionFailResponse{}

	if successResponse.Data.Code != 0 {
		err := json.Unmarshal([]byte(message), &failResponse)
		if err != nil {
			applogger.Error("Handle subscription message error: %s", err)
		}
		if wsc.responseHandler != nil {
			wsc.responseHandler(failResponse)
		}
		return
	}

	if wsc.responseHandler != nil {
		wsc.responseHandler(successResponse)
	}

}

func (wsc *WebSocketClientBase) handleDepthMessage(message string) {
	if strings.Contains(message, "snapshot") {
		snapshotDepthResponse := ws.DepthSnapshotResponse{}
		err := json.Unmarshal([]byte(message), &snapshotDepthResponse)
		if err != nil {
			applogger.Error("Handle snapshot depth message error: %s", err)
		}
		if wsc.responseHandler != nil {
			wsc.responseHandler(snapshotDepthResponse)
		}
	} else if strings.Contains(message, "update") {
		updateDepthResponse := ws.DepthUpdateResponse{}
		err := json.Unmarshal([]byte(message), &updateDepthResponse)
		if err != nil {
			applogger.Error("Handle update depth message error: %s", err)
		}
		if wsc.responseHandler != nil {
			wsc.responseHandler(updateDepthResponse)
		}
	}

}

func (wsc *WebSocketClientBase) handleOrderBookMessage(message string) {
	orderBookResponse := ws.OrderBookResponse{}
	err := json.Unmarshal([]byte(message), &orderBookResponse)
	if err != nil {
		applogger.Error("Handle orderbook message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(orderBookResponse)
	}

}

func (wsc *WebSocketClientBase) handleDepth1Message(message string) {
	depth1Response := ws.Depth1Response{}
	err := json.Unmarshal([]byte(message), &depth1Response)
	if err != nil {
		applogger.Error("Handle depth1 message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(depth1Response)
	}
}

func (wsc *WebSocketClientBase) handleTickerMessage(message string) {
	tickerResponse := ws.TickerResponse{}
	err := json.Unmarshal([]byte(message), &tickerResponse)
	if err != nil {
		applogger.Error("Handle ticker message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(tickerResponse)
	}
}

func (wsc *WebSocketClientBase) handleKlineMessage(message string) {
	klineResponse := ws.KlineResponse{}
	err := json.Unmarshal([]byte(message), &klineResponse)
	if err != nil {
		applogger.Error("Handle kline message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(klineResponse)
	}
}

func (wsc *WebSocketClientBase) handleTradeMessage(message string) {
	tradeResponse := ws.TradeResponse{}
	err := json.Unmarshal([]byte(message), &tradeResponse)
	if err != nil {
		applogger.Error("Handle trade message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(tradeResponse)
	}
}

func (wsc *WebSocketClientBase) handleMarketTradeMessage(message string) {
	marketTradeResponse := ws.MarketTradeResponse{}
	err := json.Unmarshal([]byte(message), &marketTradeResponse)
	if err != nil {
		applogger.Error("Handle market trade message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(marketTradeResponse)
	}
}

func (wsc *WebSocketClientBase) handleIndexPriceMessage(message string) {
	indexPriceResponse := ws.IndexPriceResponse{}
	err := json.Unmarshal([]byte(message), &indexPriceResponse)
	if err != nil {
		applogger.Error("Handle index price message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(indexPriceResponse)
	}
}

func (wsc *WebSocketClientBase) handleMarkPriceMessage(message string) {
	markPriceResponse := ws.MarkPriceResponse{}
	err := json.Unmarshal([]byte(message), &markPriceResponse)
	if err != nil {
		applogger.Error("Handle mark price message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(markPriceResponse)
	}
}

func (wsc *WebSocketClientBase) handleAccountMessage(message string) {
	accountResponse := ws.AccountResponse{}
	err := json.Unmarshal([]byte(message), &accountResponse)
	if err != nil {
		applogger.Error("Handle account message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(accountResponse)
	}
}

func (wsc *WebSocketClientBase) handlePositionMessage(message string) {
	positionResponse := ws.PositionResponse{}
	err := json.Unmarshal([]byte(message), &positionResponse)
	if err != nil {
		applogger.Error("Handle position message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(positionResponse)
	}
}

func (wsc *WebSocketClientBase) handleOrderMessage(message string) {
	orderResponse := ws.OrderResponse{}
	err := json.Unmarshal([]byte(message), &orderResponse)
	if err != nil {
		applogger.Error("Handle order message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(orderResponse)
	}
}

func (wsc *WebSocketClientBase) handleUserTradeMessage(message string) {
	userTradeResponse := ws.UserTradeResponse{}
	err := json.Unmarshal([]byte(message), &userTradeResponse)
	if err != nil {
		applogger.Error("Handle user trade message error: %s", err)
	}
	if wsc.responseHandler != nil {
		wsc.responseHandler(userTradeResponse)
	}
}
