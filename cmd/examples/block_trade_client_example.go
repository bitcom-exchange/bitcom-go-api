package examples

import (
	"github.com/bitcom-exchange/bitcom-go-api/config"
	"github.com/bitcom-exchange/bitcom-go-api/constant"
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/restclient"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model"
	"github.com/tidwall/pretty"
)

func NewParadigmBlockOrdersExample() {

	blockTradeClient := new(restclient.BlockTradeClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["bt_source"] = constant.ParadigmBlockTradeSource
	paramMap["label"] = "uuid-0"
	paramMap["role"] = constant.MarketMaker
	//paramMap["counterparty"] = config.User2Uid

	tradeMapOne := map[string]interface{}{
		"instrument_id": "BTC-11SEP20-10250-C",
		"price":         "0.02",
		"qty":           "20",
		"side":          constant.SellSide,
	}

	tradeMapTwo := map[string]interface{}{
		"instrument_id": "BTC-11SEP20-10250-C",
		"price":         "0.02",
		"qty":           "10",
		"side":          constant.SellSide,
	}

	tradesList := []map[string]interface{}{tradeMapOne, tradeMapTwo}

	paramMap["trades"] = tradesList

	resp, err := blockTradeClient.NewParadigmBlockOrders(paramMap)
	if err != nil {
		applogger.Error("Place block trade order error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Place block trade order: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func QueryParadigmBlockOrdersExample() {

	blockTradeClient := new(restclient.BlockTradeClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["bt_source"] = constant.ParadigmBlockTradeSource
	paramMap["label"] = "uuid-0"
	paramMap["currency"] = constant.BTC

	resp, err := blockTradeClient.QueryParadigmBlockOrders(paramMap)
	if err != nil {
		applogger.Error("Query block trade order error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Query block trade order: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func QueryParadigmBlockOrdersByPlatformExample() {

	blockTradeClient := new(restclient.BlockTradeClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["bt_source"] = constant.ParadigmBlockTradeSource
	paramMap["currency"] = constant.BTC
	paramMap["maker"] = config.User1Uid

	resp, err := blockTradeClient.QueryParadigmBlockOrdersByPlatform(paramMap)
	if err != nil {
		applogger.Error("Query block trade orders by platform error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Query block trade orders by platform: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}
