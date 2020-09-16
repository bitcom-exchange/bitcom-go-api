package restclient

import (
	"encoding/json"
	"errors"
	"github.com/bitcom-exchange/bitcom-go-api/constant"
	"github.com/bitcom-exchange/bitcom-go-api/internal"
	"github.com/bitcom-exchange/bitcom-go-api/internal/requestbuilder"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/account"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"
)

type AccountClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

func (c *AccountClient) Init(host string, accessKey string, secretKey string) *AccountClient {
	c.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(host, accessKey, secretKey)
	return c
}

func (c *AccountClient) GetAccounts(currency string) (*account.GetAccountResponse, error) {
	paramMap := make(map[string]interface{})
	paramMap["currency"] = currency
	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetAccountsUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, c.privateUrlBuilder.GetAccessKey())
	result := &account.GetAccountResponse{}

	if getErr != nil {
		return nil, getErr
	}

	jsonErr := json.Unmarshal([]byte(getResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *AccountClient) GetPositions(paramMap map[string]interface{}) (*account.GetPositionsResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetPositionsUrl, paramMap)

	getResp, getErr := internal.HttpGet(url, c.privateUrlBuilder.GetAccessKey())
	if getErr != nil {
		return nil, getErr
	}
	baseResult := &base.RestBaseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), baseResult)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if baseResult.Code == 0 {
		result := &account.GetPositionsResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *AccountClient) GetTransactionLogs(paramMap map[string]interface{}) (*account.GetTransactionLogsResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}
	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetTransactionLogsUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, c.privateUrlBuilder.GetAccessKey())
	if getErr != nil {
		return nil, getErr
	}
	baseResult := &base.RestBaseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), baseResult)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if baseResult.Code == 0 {
		result := &account.GetTransactionLogsResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *AccountClient) GetUserDeliveries(paramMap map[string]interface{}) (*account.GetUserDeliveriesResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetUserDeliveriesUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, c.privateUrlBuilder.GetAccessKey())
	if getErr != nil {
		return nil, getErr
	}
	baseResult := &base.RestBaseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), baseResult)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if baseResult.Code == 0 {
		result := &account.GetUserDeliveriesResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *AccountClient) GetUserSettlements(paramMap map[string]interface{}) (*account.GetUserSettlementsResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetUserSettlementsUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, c.privateUrlBuilder.GetAccessKey())
	if getErr != nil {
		return nil, getErr
	}
	baseResult := &base.RestBaseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), baseResult)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if baseResult.Code == 0 {
		result := &account.GetUserSettlementsResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}
