package examples

import (
	"github.com/bitcom-exchange/bitcom-go-api/config"
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/restclient"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model"
	"github.com/tidwall/pretty"
)

func PlaceNewOrderExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["qty"] = "1000"
	paramMap["side"] = "sell"
	paramMap["price"] = "10997.00"
	paramMap["order_type"] = "limit"

	resp, err := orderClient.NewOrder(paramMap)
	if err != nil {
		applogger.Error("Place new order error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Place new order: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func PlaceNewBatchOrderExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	orderMapOne := make(map[string]interface{})

	orderMapOne["instrument_id"] = "BTC-PERPETUAL"
	orderMapOne["qty"] = "1000"
	orderMapOne["side"] = "buy"
	orderMapOne["price"] = "10380.00"
	orderMapOne["order_type"] = "limit"

	orderMapTwo := make(map[string]interface{})

	orderMapTwo["instrument_id"] = "BTC-PERPETUAL"
	orderMapTwo["qty"] = "1000"
	orderMapTwo["side"] = "sell"
	orderMapTwo["price"] = "10410.00"
	orderMapTwo["order_type"] = "limit"

	orderSlice := []map[string]interface{}{orderMapOne, orderMapTwo}

	paramMap := map[string]interface{}{
		"orders_data": orderSlice,
	}

	resp, err := orderClient.NewBatchOrders(paramMap)
	if err != nil {
		applogger.Error("Place batch new orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Place batch new orders: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func CancelOrderExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["order_id"] = "2378090"

	resp, err := orderClient.CancelOrders(paramMap)
	if err != nil {
		applogger.Error("Cancel orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Cancel orders: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func AmendOrderExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["order_id"] = "1373134"
	paramMap["price"] = "10600.00"

	resp, err := orderClient.AmendOrder(paramMap)
	if err != nil {
		applogger.Error("Amend order error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Amend order: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func AmendBatchOrdersExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	orderMapOne := make(map[string]interface{})

	orderMapOne["order_id"] = "1373574"
	orderMapOne["price"] = "10701.00"

	orderMapTwo := make(map[string]interface{})

	orderMapTwo["order_id"] = "1373134"
	orderMapTwo["price"] = "10601.00"

	orderSlice := []map[string]interface{}{orderMapOne, orderMapTwo}

	paramMap := map[string]interface{}{
		"orders_data": orderSlice,
	}

	resp, err := orderClient.AmendBatchOrders(paramMap)
	if err != nil {
		applogger.Error("Amend batch orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Amend batch orders: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func ClosePositionsExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["order_type"] = "limit"
	paramMap["price"] = "10286.50"

	resp, err := orderClient.ClosePositions(paramMap)
	if err != nil {
		applogger.Error("Close positions error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Close positions: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetOpenOrdersExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "future"

	resp, err := orderClient.GetOpenOrders(paramMap)
	if err != nil {
		applogger.Error("Get open orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get open orders: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetOrdersExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "future"
	paramMap["instrument_id"] = "BTC-PERPETUAL"

	resp, err := orderClient.GetOrders(paramMap)
	if err != nil {
		applogger.Error("Get orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get orders: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetStopOrdersExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["instrument_id"] = "BTC-PERPETUAL"

	resp, err := orderClient.GetStopOrders(paramMap)
	if err != nil {
		applogger.Error("Get stop orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get stop orders: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetUserTradesExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["instrument_id"] = "BTC-PERPETUAL"

	resp, err := orderClient.GetUserTrades(paramMap)
	if err != nil {
		applogger.Error("Get user trades error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get user trades: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetEstMarginsExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["price"] = "10318.00"
	paramMap["qty"] = "2000"

	resp, err := orderClient.GetEstMargins(paramMap)
	if err != nil {
		applogger.Error("Get estimated margins error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get estimated margins: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}
