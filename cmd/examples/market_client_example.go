package examples

import (
	"github.com/bitcom-exchange/bitcom-go-api/config"
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/restclient"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model"
	"github.com/tidwall/pretty"
)

func GetIndexExample() {
	systemClient := new(restclient.MarketClient).Init(config.User1Host)
	resp, err := systemClient.GetIndex("BTC")
	if err != nil {
		applogger.Error("Get index error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get index: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetInstrumentsExample() {
	systemClient := new(restclient.MarketClient).Init(config.User1Host)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "future"
	paramMap["active"] = true

	resp, err := systemClient.GetInstruments(paramMap)
	if err != nil {
		applogger.Error("Get instruments error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get instruments: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetTickerExample() {
	systemClient := new(restclient.MarketClient).Init(config.User1Host)
	resp, err := systemClient.GetTicker("BTC-10SEP20-10125-C")
	if err != nil {
		applogger.Error("Get ticker error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get ticker: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetOrderBookExample() {
	systemClient := new(restclient.MarketClient).Init(config.User1Host)

	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["level"] = 5

	resp, err := systemClient.GetOrderBook(paramMap)
	if err != nil {
		applogger.Error("Get order book error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get order book: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetMarketTradeExample() {
	systemClient := new(restclient.MarketClient).Init(config.User1Host)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "option"
	paramMap["option_type"] = "call"
	paramMap["instrument_id"] = ""
	paramMap["offset"] = 1
	paramMap["limit"] = 5

	resp, err := systemClient.GetMarketTrade(paramMap)
	if err != nil {
		applogger.Error("Get market trade error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get market trade: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetKlinesExample() {
	systemClient := new(restclient.MarketClient).Init(config.User1Host)

	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["start_time"] = 1599654319564
	paramMap["end_time"] = 1599654430564
	paramMap["timeframe_min"] = "1"

	resp, err := systemClient.GetKlines(paramMap)
	if err != nil {
		applogger.Error("Get K-line error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get K-line: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetDeliveryInfoExample() {
	systemClient := new(restclient.MarketClient).Init(config.User1Host)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["offset"] = 1
	paramMap["limit"] = 5

	resp, err := systemClient.GetDeliveryInfo(paramMap)
	if err != nil {
		applogger.Error("Get daily delivery price info error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get daily delivery price info: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetMarketSummaryExample() {
	systemClient := new(restclient.MarketClient).Init(config.User1Host)
	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "future"
	paramMap["instrument_id"] = "BTC-PERPETUAL"

	resp, err := systemClient.GetMarketSummary(paramMap)
	if err != nil {
		applogger.Error("Get market summary error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get market summary: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetFundingRateExample() {
	systemClient := new(restclient.MarketClient).Init(config.User1Host)
	resp, err := systemClient.GetFundingRate("BTC-PERPETUAL")
	if err != nil {
		applogger.Error("Get funding rate error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get funding rate: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}
