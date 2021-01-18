package examples

import (
	"github.com/bitcom-exchange/bitcom-go-api/config"
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/restclient"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model"
	"github.com/tidwall/pretty"
)

func GetIndexExample() {
	marketClient := new(restclient.MarketClient).Init(config.User1Host)
	resp, err := marketClient.GetIndex("BTC")
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
	marketClient := new(restclient.MarketClient).Init(config.User1Host)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "future"
	paramMap["active"] = true

	resp, err := marketClient.GetInstruments(paramMap)
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
	marketClient := new(restclient.MarketClient).Init(config.User1Host)
	resp, err := marketClient.GetTicker("BTC-10SEP20-10125-C")
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
	marketClient := new(restclient.MarketClient).Init(config.User1Host)

	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["level"] = 5

	resp, err := marketClient.GetOrderBook(paramMap)
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
	marketClient := new(restclient.MarketClient).Init(config.User1Host)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "option"
	paramMap["option_type"] = "call"
	paramMap["instrument_id"] = ""
	paramMap["offset"] = 1
	paramMap["limit"] = 5

	resp, err := marketClient.GetMarketTrade(paramMap)
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
	marketClient := new(restclient.MarketClient).Init(config.User1Host)

	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["start_time"] = 1599654319564
	paramMap["end_time"] = 1599654430564
	paramMap["timeframe_min"] = "1"

	resp, err := marketClient.GetKlines(paramMap)
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
	marketClient := new(restclient.MarketClient).Init(config.User1Host)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["offset"] = 1
	paramMap["limit"] = 5

	resp, err := marketClient.GetDeliveryInfo(paramMap)
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
	marketClient := new(restclient.MarketClient).Init(config.User1Host)
	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "future"
	paramMap["instrument_id"] = "BTC-PERPETUAL"

	resp, err := marketClient.GetMarketSummary(paramMap)
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
	marketClient := new(restclient.MarketClient).Init(config.User1Host)
	resp, err := marketClient.GetFundingRate("BTC-PERPETUAL")
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

func GetFundingRateHistoryExample() {
	marketClient := new(restclient.MarketClient).Init(config.User1Host)
	paramMap := map[string]interface{}{
		"instrument_id": "BTC-PERPETUAL",
		"start_time":    1610784000000,
		"end_time":      1610870400000,
		"history_type":  "8H",
	}

	resp, err := marketClient.GetFundingRateHistory(paramMap)
	if err != nil {
		applogger.Error("Get funding rate history error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get funding rate history: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetTotalVolumeExample() {
	marketClient := new(restclient.MarketClient).Init(config.User1Host)

	resp, err := marketClient.GetTotalVolume()
	if err != nil {
		applogger.Error("Get total volume error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get total volume: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}
