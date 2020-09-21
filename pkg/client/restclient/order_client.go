package restclient

import (
	"encoding/json"
	"errors"
	"github.com/bitcom-exchange/bitcom-go-api/constant"
	"github.com/bitcom-exchange/bitcom-go-api/internal"
	"github.com/bitcom-exchange/bitcom-go-api/internal/requestbuilder"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/order"
)

type OrderClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

func (c *OrderClient) Init(host string, accessKey string, secretKey string) *OrderClient {
	c.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(host, accessKey, secretKey)
	return c
}

func (c *OrderClient) NewOrder(paramMap map[string]interface{}) (*order.NewOrderResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url, postBody, err := c.privateUrlBuilder.BuildWithBody(constant.V1NewOrderUrl, paramMap)
	if err != nil {
		return nil, err
	}

	postResp, postErr := internal.HttpPost(url, postBody, c.privateUrlBuilder.GetAccessKey())
	if postErr != nil {
		return nil, postErr
	}
	result := &order.NewOrderResponse{}

	jsonErr := json.Unmarshal([]byte(postResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(postResp)
}

func (c *OrderClient) NewBatchOrders(paramMap map[string]interface{}) (*order.NewBatchOrdersResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url, postBody, err := c.privateUrlBuilder.BuildWithBody(constant.V1NewBatchOrdersUrl, paramMap)
	if err != nil {
		return nil, err
	}

	postResp, postErr := internal.HttpPost(url, postBody, c.privateUrlBuilder.GetAccessKey())
	if postErr != nil {
		return nil, postErr
	}
	result := &order.NewBatchOrdersResponse{}

	jsonErr := json.Unmarshal([]byte(postResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(postResp)
}

func (c *OrderClient) CancelOrders(paramMap map[string]interface{}) (*order.NewCancelOrdersResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url, postBody, err := c.privateUrlBuilder.BuildWithBody(constant.V1CancelOrdersUrl, paramMap)
	if err != nil {
		return nil, err
	}

	postResp, postErr := internal.HttpPost(url, postBody, c.privateUrlBuilder.GetAccessKey())
	if postErr != nil {
		return nil, postErr
	}
	result := &order.NewCancelOrdersResponse{}

	jsonErr := json.Unmarshal([]byte(postResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(postResp)
}

func (c *OrderClient) AmendOrder(paramMap map[string]interface{}) (*order.AmendOrderResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url, postBody, err := c.privateUrlBuilder.BuildWithBody(constant.V1AmendOrdersUrl, paramMap)
	if err != nil {
		return nil, err
	}

	postResp, postErr := internal.HttpPost(url, postBody, c.privateUrlBuilder.GetAccessKey())
	if postErr != nil {
		return nil, postErr
	}
	result := &order.AmendOrderResponse{}

	jsonErr := json.Unmarshal([]byte(postResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(postResp)
}

func (c *OrderClient) AmendBatchOrders(paramMap map[string]interface{}) (*order.AmendBatchOrdersResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url, postBody, err := c.privateUrlBuilder.BuildWithBody(constant.V1AmendBatchOrdersUrl, paramMap)
	if err != nil {
		return nil, err
	}

	postResp, postErr := internal.HttpPost(url, postBody, c.privateUrlBuilder.GetAccessKey())
	if postErr != nil {
		return nil, postErr
	}
	result := &order.AmendBatchOrdersResponse{}

	jsonErr := json.Unmarshal([]byte(postResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(postResp)
}

func (c *OrderClient) ClosePositions(paramMap map[string]interface{}) (*order.ClosePositionsResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url, postBody, err := c.privateUrlBuilder.BuildWithBody(constant.V1ClosePositionsUrl, paramMap)
	if err != nil {
		return nil, err
	}

	postResp, postErr := internal.HttpPost(url, postBody, c.privateUrlBuilder.GetAccessKey())
	if postErr != nil {
		return nil, postErr
	}
	result := &order.ClosePositionsResponse{}

	jsonErr := json.Unmarshal([]byte(postResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(postResp)
}

func (c *OrderClient) GetOpenOrders(paramMap map[string]interface{}) (*order.GetOpenOrdersResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetOpenOrdersUrl, paramMap)

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
		result := &order.GetOpenOrdersResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *OrderClient) GetOrders(paramMap map[string]interface{}) (*order.GetOrdersResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetOrdersUrl, paramMap)

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
		result := &order.GetOrdersResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *OrderClient) GetStopOrders(paramMap map[string]interface{}) (*order.GetStopOrdersResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetStopOrdersUrl, paramMap)

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
		result := &order.GetStopOrdersResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *OrderClient) GetUserTrades(paramMap map[string]interface{}) (*order.GetUserTradesResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetUserTradesUrl, paramMap)

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
		result := &order.GetUserTradesResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *OrderClient) GetEstMargins(paramMap map[string]interface{}) (*order.GetEstMarginsResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1GetEstMarginsUrl, paramMap)

	getResp, getErr := internal.HttpGet(url, c.privateUrlBuilder.GetAccessKey())
	if getErr != nil {
		return nil, getErr
	}
	result := &order.GetEstMarginsResponse{}

	jsonErr := json.Unmarshal([]byte(getResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(getResp)
}
