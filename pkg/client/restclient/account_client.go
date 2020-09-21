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

func (c *AccountClient) GetCodConfig(paramMap map[string]interface{}) (*account.GetCodResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetCodConfigUrl, paramMap)

	getResp, getErr := internal.HttpGet(url, c.privateUrlBuilder.GetAccessKey())
	if getErr != nil {
		return nil, getErr
	}
	result := &account.GetCodResponse{}

	jsonErr := json.Unmarshal([]byte(getResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *AccountClient) ConfigCod(paramMap map[string]interface{}) (*account.PostCodResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url, postBody, err := c.privateUrlBuilder.BuildWithBody(constant.V1ConfigCodUrl, paramMap)
	if err != nil {
		return nil, err
	}

	postResp, postErr := internal.HttpPost(url, postBody, c.privateUrlBuilder.GetAccessKey())
	if postErr != nil {
		return nil, postErr
	}
	result := &account.PostCodResponse{}

	jsonErr := json.Unmarshal([]byte(postResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(postResp)
}

func (c *AccountClient) GetMmpState(paramMap map[string]interface{}) (*account.GetMmpResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetMmpStateUrl, paramMap)

	getResp, getErr := internal.HttpGet(url, c.privateUrlBuilder.GetAccessKey())
	if getErr != nil {
		return nil, getErr
	}
	result := &account.GetMmpResponse{}

	jsonErr := json.Unmarshal([]byte(getResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *AccountClient) UpdateMmpConfig(paramMap map[string]interface{}) (*account.UpdateMmpResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url, postBody, err := c.privateUrlBuilder.BuildWithBody(constant.V1UpdateMmpConfigUrl, paramMap)
	if err != nil {
		return nil, err
	}

	postResp, postErr := internal.HttpPost(url, postBody, c.privateUrlBuilder.GetAccessKey())
	if postErr != nil {
		return nil, postErr
	}
	result := &account.UpdateMmpResponse{}

	jsonErr := json.Unmarshal([]byte(postResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(postResp)
}

func (c *AccountClient) ResetMmpState(paramMap map[string]interface{}) (*account.ResetMmpStateResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url, postBody, err := c.privateUrlBuilder.BuildWithBody(constant.V1UpdateMmpConfigUrl, paramMap)
	if err != nil {
		return nil, err
	}

	postResp, postErr := internal.HttpPost(url, postBody, c.privateUrlBuilder.GetAccessKey())
	if postErr != nil {
		return nil, postErr
	}
	result := &account.ResetMmpStateResponse{}

	jsonErr := json.Unmarshal([]byte(postResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(postResp)
}
