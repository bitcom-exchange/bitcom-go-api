package examples

import (
	"fmt"
	"github.com/bitcom-exchange/bitcom-go-api/config"
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/restclient"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/wsclient"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/ws"
)

func responseHandlerExample(resp interface{}) {
	switch t := resp.(type) {
	case ws.SubscriptionSuccessResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive subscription success response: %s", respJson)
	case ws.SubscriptionFailResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive subscription fail response: %s", respJson)
	case ws.DepthSnapshotResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive depth snapshot response: %s", respJson)
	case ws.DepthUpdateResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive depth update response: %s", respJson)
	case ws.OrderBookResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive orderbook response: %s", respJson)
	case ws.Depth1Response:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive depth1 response: %s", respJson)
	case ws.TickerResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive ticker response: %s", respJson)
	case ws.KlineResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive kline response: %s", respJson)
	case ws.TradeResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive trade response: %s", respJson)
	case ws.MarketTradeResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive market trade response: %s", respJson)
	case ws.IndexPriceResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive index price response: %s", respJson)
	case ws.MarkPriceResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive mark price response: %s", respJson)
	case ws.AccountResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive account response: %s", respJson)
	case ws.PositionResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive position response: %s", respJson)
	case ws.OrderResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive order response: %s", respJson)
	case ws.UserTradeResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive user trades response: %s", respJson)
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}

func getWsAuthToken() string {
	wsAuthClient := new(restclient.WsAuthClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)
	resp, err := wsAuthClient.GetWsAuthToken()
	if err != nil {
		applogger.Error("Get auth token error: %s", err)
		return ""
	} else {
		return resp.Data.Token
	}
}

func PublicSubscribeExample() {

	client := new(wsclient.PublicWebsocketClient).Init(config.WsHost)

	paramMap := map[string]interface{}{
		"type":     "subscribe",
		"channels": []string{"mark_price"},
		//"currencies": []string{"BTC"},
		//"categories": []string{"option"},
		"instruments": []string{"BTC-PERPETUAL", "BTC-15SEP20-10625-C"},
		"interval":    "100ms",
	}

	client.SetHandler(
		func() {
			client.Subscribe(paramMap)
		},
		responseHandlerExample,
	)

	client.Connect(false)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	paramMap["type"] = "unsubscribe"
	client.UnSubscribe(paramMap)

	client.Close()
	applogger.Info("Client closed")
}

func PrivateSubscribeExample() {
	token := getWsAuthToken()

	client := new(wsclient.PrivateWebsocketClient).Init(config.WsHost, token)

	paramMap := map[string]interface{}{
		"type":        "subscribe",
		"instruments": []string{"BTC-PERPETUAL"},
		"channels":    []string{"order"},
		"currencies":  []string{"BTC"},
		"categories":  []string{"future"},
		"interval":    "100ms",
	}

	client.SetHandler(
		func() {
			client.Subscribe(paramMap)
		},
		responseHandlerExample,
	)

	client.Connect(false)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	paramMap["type"] = "unsubscribe"
	client.UnSubscribe(paramMap)

	client.Close()
	applogger.Info("Client closed")
}
