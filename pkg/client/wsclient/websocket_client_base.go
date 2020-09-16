package wsclient

import (
	"encoding/json"
	"fmt"
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/ws"
	"github.com/gorilla/websocket"
	"strings"
	"sync"
	"time"
)

const (
	TimerIntervalSecond = 5
	ReconnectWaitSecond = 60

	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	// TODO: Set pong wait to a long time to avoid network problems
	pongWait = 10 * time.Minute

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = 54 * time.Second // (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 2 * 1024 * 1024

	writeQueueSize = 512

	path = "/"
)

type ConnectedHandler func()

type MessageHandler func(message string) (interface{}, error)

type ResponseHandler func(response interface{})

type WebSocketClientBase struct {
	host              string
	token             string
	conn              *websocket.Conn
	connectedHandler  ConnectedHandler
	responseHandler   ResponseHandler
	stopReadChannel   chan int
	stopTickerChannel chan int
	ticker            *time.Ticker
	lastReceivedTime  time.Time
	sendMutex         *sync.Mutex
}

func (p *WebSocketClientBase) Init(host, token string) *WebSocketClientBase {
	p.host = host
	p.token = token
	p.stopReadChannel = make(chan int, 1)
	p.stopTickerChannel = make(chan int, 1)
	p.sendMutex = &sync.Mutex{}

	return p
}

func (p *WebSocketClientBase) BuildRequestBody(paramMap map[string]interface{}) (string, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	if p.token != "" {
		paramMap["token"] = p.token
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

func (p *WebSocketClientBase) SetHandler(connHandler ConnectedHandler, repHandler ResponseHandler) {
	p.connectedHandler = connHandler
	p.responseHandler = repHandler
}

func (p *WebSocketClientBase) Connect(autoConnect bool) {
	p.connectWebSocket()

	if autoConnect {
		p.startTicker()
	}
}

func (p *WebSocketClientBase) Send(data string) {
	if p.conn == nil {
		applogger.Error("WebSocket sent error: no connection available")
		return
	}

	p.sendMutex.Lock()
	err := p.conn.WriteMessage(websocket.TextMessage, []byte(data))
	p.sendMutex.Unlock()

	if err != nil {
		applogger.Error("WebSocket sent error: data=%s, error=%s", data, err)
	}
}

func (p *WebSocketClientBase) Close() {
	//p.stopTicker()
	p.stopReadLoop()
	p.disconnectWebSocket()
}

func (p *WebSocketClientBase) connectWebSocket() {
	var err error
	url := fmt.Sprintf("%s%s", p.host, path)
	applogger.Debug("WebSocket connecting...")
	p.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		applogger.Error("WebSocket connected error: %s", err)
		return
	}
	applogger.Info("WebSocket connected")

	p.conn.SetPingHandler(nil)
	p.conn.SetReadLimit(maxMessageSize)

	if err := p.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		applogger.Error("WebSocket connected error: %s", err)
		return
	}

	p.conn.SetPongHandler(func(string) error {
		return p.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	p.startReadLoop()

	if p.connectedHandler != nil {
		p.connectedHandler()
	}
}

func (p *WebSocketClientBase) disconnectWebSocket() {
	if p.conn == nil {
		return
	}

	applogger.Debug("WebSocket disconnecting...")

	time.Sleep(time.Second)

	err := p.conn.Close()
	if err != nil {
		applogger.Error("WebSocket disconnect error: %s", err)
		return
	}

	applogger.Info("WebSocket disconnected")
}

func (p *WebSocketClientBase) startTicker() {
	p.ticker = time.NewTicker(TimerIntervalSecond * time.Second)
	p.lastReceivedTime = time.Now()

	go p.tickerLoop()
}

func (p *WebSocketClientBase) stopTicker() {
	p.ticker.Stop()
	p.stopTickerChannel <- 1
}

func (p *WebSocketClientBase) tickerLoop() {
	applogger.Debug("tickerLoop started")
	for {
		select {
		// Receive data from stopChannel
		case <-p.stopTickerChannel:
			applogger.Debug("tickerLoop stopped")
			return

		// Receive tick from tickChannel
		case <-p.ticker.C:
			elapsedSecond := time.Now().Sub(p.lastReceivedTime).Seconds()
			applogger.Debug("WebSocket received data %f sec ago", elapsedSecond)

			if elapsedSecond > ReconnectWaitSecond {
				applogger.Info("WebSocket reconnect...")
				p.disconnectWebSocket()
				p.connectWebSocket()
			}
		}
	}
}

func (p *WebSocketClientBase) startReadLoop() {
	go p.readLoop()
}

func (p *WebSocketClientBase) stopReadLoop() {
	p.stopReadChannel <- 1
}

func (p *WebSocketClientBase) readLoop() {
	applogger.Debug("readLoop started")
	for {
		select {
		case <-p.stopReadChannel:
			applogger.Debug("readLoop stopped")
			return

		default:
			if p.conn == nil {
				applogger.Error("Read error: no connection available")
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			// TODO: If no message receives, the function will be blocked here
			msgType, buf, err := p.conn.ReadMessage()
			if err != nil {
				applogger.Error("Read error: %s", err)
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}
			p.lastReceivedTime = time.Now()

			switch msgType {
			case websocket.PingMessage:
				applogger.Info("Read ping message: %s", string(buf))
			case websocket.TextMessage:
				message := string(buf)
				p.handleTextMessage(message)
			}

		}
	}
}

func (p *WebSocketClientBase) handleTextMessage(message string) {
	switch {
	case strings.Contains(message, "\"channel\":\"subscription\""):
		p.handleSubscriptionMessage(message)
	case strings.Contains(message, "\"channel\":\"depth\"") && !strings.Contains(message, "depth1"):
		p.handleDepthMessage(message)
	case strings.Contains(message, "\"channel\":\"order_book\""):
		p.handleOrderBookMessage(message)
	case strings.Contains(message, "\"channel\":\"depth1\""):
		p.handleDepth1Message(message)
	case strings.Contains(message, "\"channel\":\"ticker\""):
		p.handleTickerMessage(message)
	case strings.Contains(message, "\"channel\":\"ticker\""):
		p.handleKlineMessage(message)
	case strings.Contains(message, "\"channel\":\"trade\""):
		p.handleTradeMessage(message)
	case strings.Contains(message, "\"channel\":\"market_trade\""):
		p.handleMarketTradeMessage(message)
	case strings.Contains(message, "\"channel\":\"index_price\""):
		p.handleIndexPriceMessage(message)
	case strings.Contains(message, "\"channel\":\"mark_price\""):
		p.handleMarkPriceMessage(message)
	case strings.Contains(message, "\"channel\":\"account\""):
		p.handleAccountMessage(message)
	case strings.Contains(message, "\"channel\":\"position\""):
		p.handlePositionMessage(message)
	case strings.Contains(message, "\"channel\":\"order\"") && !strings.Contains(message, "order_book"):
		p.handleOrderMessage(message)
	case strings.Contains(message, "\"channel\":\"user_trade\""):
		p.handleUserTradeMessage(message)
	}
}

func (p *WebSocketClientBase) handleSubscriptionMessage(message string) {
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
		if p.responseHandler != nil {
			p.responseHandler(failResponse)
		}
		return
	}

	if p.responseHandler != nil {
		p.responseHandler(successResponse)
	}

}

func (p *WebSocketClientBase) handleDepthMessage(message string) {
	if strings.Contains(message, "snapshot") {
		snapshotDepthResponse := ws.DepthSnapshotResponse{}
		err := json.Unmarshal([]byte(message), &snapshotDepthResponse)
		if err != nil {
			applogger.Error("Handle snapshot depth message error: %s", err)
		}
		if p.responseHandler != nil {
			p.responseHandler(snapshotDepthResponse)
		}
	} else if strings.Contains(message, "update") {
		updateDepthResponse := ws.DepthUpdateResponse{}
		err := json.Unmarshal([]byte(message), &updateDepthResponse)
		if err != nil {
			applogger.Error("Handle update depth message error: %s", err)
		}
		if p.responseHandler != nil {
			p.responseHandler(updateDepthResponse)
		}
	}

}

func (p *WebSocketClientBase) handleOrderBookMessage(message string) {
	orderBookResponse := ws.OrderBookResponse{}
	err := json.Unmarshal([]byte(message), &orderBookResponse)
	if err != nil {
		applogger.Error("Handle orderbook message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(orderBookResponse)
	}

}

func (p *WebSocketClientBase) handleDepth1Message(message string) {
	depth1Response := ws.Depth1Response{}
	err := json.Unmarshal([]byte(message), &depth1Response)
	if err != nil {
		applogger.Error("Handle depth1 message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(depth1Response)
	}
}

func (p *WebSocketClientBase) handleTickerMessage(message string) {
	tickerResponse := ws.TickerResponse{}
	err := json.Unmarshal([]byte(message), &tickerResponse)
	if err != nil {
		applogger.Error("Handle ticker message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(tickerResponse)
	}
}

func (p *WebSocketClientBase) handleKlineMessage(message string) {
	klineResponse := ws.KlineResponse{}
	err := json.Unmarshal([]byte(message), &klineResponse)
	if err != nil {
		applogger.Error("Handle kline message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(klineResponse)
	}
}

func (p *WebSocketClientBase) handleTradeMessage(message string) {
	tradeResponse := ws.TradeResponse{}
	err := json.Unmarshal([]byte(message), &tradeResponse)
	if err != nil {
		applogger.Error("Handle trade message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(tradeResponse)
	}
}

func (p *WebSocketClientBase) handleMarketTradeMessage(message string) {
	marketTradeResponse := ws.MarketTradeResponse{}
	err := json.Unmarshal([]byte(message), &marketTradeResponse)
	if err != nil {
		applogger.Error("Handle market trade message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(marketTradeResponse)
	}
}

func (p *WebSocketClientBase) handleIndexPriceMessage(message string) {
	indexPriceResponse := ws.IndexPriceResponse{}
	err := json.Unmarshal([]byte(message), &indexPriceResponse)
	if err != nil {
		applogger.Error("Handle index price message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(indexPriceResponse)
	}
}

func (p *WebSocketClientBase) handleMarkPriceMessage(message string) {
	markPriceResponse := ws.MarkPriceResponse{}
	err := json.Unmarshal([]byte(message), &markPriceResponse)
	if err != nil {
		applogger.Error("Handle mark price message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(markPriceResponse)
	}
}

func (p *WebSocketClientBase) handleAccountMessage(message string) {
	accountResponse := ws.AccountResponse{}
	err := json.Unmarshal([]byte(message), &accountResponse)
	if err != nil {
		applogger.Error("Handle account message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(accountResponse)
	}
}

func (p *WebSocketClientBase) handlePositionMessage(message string) {
	positionResponse := ws.PositionResponse{}
	err := json.Unmarshal([]byte(message), &positionResponse)
	if err != nil {
		applogger.Error("Handle position message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(positionResponse)
	}
}

func (p *WebSocketClientBase) handleOrderMessage(message string) {
	orderResponse := ws.OrderResponse{}
	err := json.Unmarshal([]byte(message), &orderResponse)
	if err != nil {
		applogger.Error("Handle order message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(orderResponse)
	}
}

func (p *WebSocketClientBase) handleUserTradeMessage(message string) {
	userTradeResponse := ws.UserTradeResponse{}
	err := json.Unmarshal([]byte(message), &userTradeResponse)
	if err != nil {
		applogger.Error("Handle user trade message error: %s", err)
	}
	if p.responseHandler != nil {
		p.responseHandler(userTradeResponse)
	}
}
