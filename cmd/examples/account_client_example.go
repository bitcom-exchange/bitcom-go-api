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

func GetCodConfigExample() {
	accountClient := new(restclient.AccountClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"

	resp, err := accountClient.GetCodConfig(paramMap)
	if err != nil {
		applogger.Error("Get cod config error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get cod config: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func ConfigCodExample() {
	accountClient := new(restclient.AccountClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["cod"] = true

	resp, err := accountClient.ConfigCod(paramMap)
	if err != nil {
		applogger.Error("Get cod config error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get cod config: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetMmpStateExample() {
	accountClient := new(restclient.AccountClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"

	resp, err := accountClient.GetMmpState(paramMap)
	if err != nil {
		applogger.Error("Get mmp state error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get mmp state: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func UpdateMmpConfigExample() {
	accountClient := new(restclient.AccountClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["window_ms"] = 20000
	paramMap["frozen_period_ms"] = 30000
	paramMap["qty_limit"] = "1000.00000000"
	paramMap["delta_limit"] = "1000.00000000"

	resp, err := accountClient.UpdateMmpConfig(paramMap)
	if err != nil {
		applogger.Error("Update mmp config error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Update mmp config: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func ResetMmpStateExample() {
	accountClient := new(restclient.AccountClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"

	resp, err := accountClient.ResetMmpState(paramMap)
	if err != nil {
		applogger.Error("Reset mmp state error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Reset mmp state: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}
