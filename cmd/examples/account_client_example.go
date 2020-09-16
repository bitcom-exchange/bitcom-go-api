package examples

import (
	"github.com/bitcom-exchange/bitcom-go-api/config"
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/restclient"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model"
	"github.com/tidwall/pretty"
)

func GetAccountsExample() {
	accountClient := new(restclient.AccountClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)
	resp, err := accountClient.GetAccounts("BTC")
	if err != nil {
		applogger.Error("Get accounts error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get accounts: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetPositionsExample() {
	accountClient := new(restclient.AccountClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "future"
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["offset"] = 1
	paramMap["limit"] = 5

	resp, err := accountClient.GetPositions(paramMap)
	if err != nil {
		applogger.Error("Get positions error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get positions: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetTransactionLogsExample() {
	accountClient := new(restclient.AccountClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["offset"] = 1
	paramMap["limit"] = 5

	resp, err := accountClient.GetTransactionLogs(paramMap)
	if err != nil {
		applogger.Error("Get transaction logs error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get transaction logs: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetUserDeliveriesExample() {
	accountClient := new(restclient.AccountClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)
	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["offset"] = 1
	paramMap["limit"] = 5

	resp, err := accountClient.GetUserDeliveries(paramMap)
	if err != nil {
		applogger.Error("Get user deliveries error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get user deliveries: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetUserSettlementsExample() {
	accountClient := new(restclient.AccountClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)
	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["offset"] = 1
	paramMap["limit"] = 5

	resp, err := accountClient.GetUserSettlements(paramMap)
	if err != nil {
		applogger.Error("Get user settlements error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get user settlements: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}
