package restclient

import (
	"encoding/json"
	"errors"
	"github.com/bitcom-exchange/bitcom-go-api/constant"
	"github.com/bitcom-exchange/bitcom-go-api/internal"
	"github.com/bitcom-exchange/bitcom-go-api/internal/requestbuilder"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/market"
)

type MarketClient struct {
	publicUrlBuilder *requestbuilder.PublicUrlBuilder
}

func (c *MarketClient) Init(host string) *MarketClient {
	c.publicUrlBuilder = new(requestbuilder.PublicUrlBuilder).Init(host)
	return c
}

func (c *MarketClient) GetIndex(currency string) (*market.GetIndexResponse, error) {
	paramMap := make(map[string]interface{})
	paramMap["currency"] = currency
	url := c.publicUrlBuilder.Build(constant.V1GetIndexUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, "")

	result := &market.GetIndexResponse{}

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

func (c *MarketClient) GetInstruments(paramMap map[string]interface{}) (*market.GetInstrumentsResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.publicUrlBuilder.Build(constant.V1GetInstrumentsUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, "")
	if getErr != nil {
		return nil, getErr
	}
	baseResult := &base.RestBaseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), baseResult)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if baseResult.Code == 0 {
		result := &market.GetInstrumentsResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *MarketClient) GetTicker(instrumentId string) (*market.GetTickerResponse, error) {
	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = instrumentId

	url := c.publicUrlBuilder.Build(constant.V1GetTickersUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, "")

	result := &market.GetTickerResponse{}

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

func (c *MarketClient) GetOrderBook(paramMap map[string]interface{}) (*market.GetOrderBookResponse, error) {

	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.publicUrlBuilder.Build(constant.V1GetOrderbooksUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, "")

	result := &market.GetOrderBookResponse{}

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

func (c *MarketClient) GetMarketTrade(paramMap map[string]interface{}) (*market.GetMarketTradeResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.publicUrlBuilder.Build(constant.V1GetMarketTradesUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, "")
	if getErr != nil {
		return nil, getErr
	}
	baseResult := &base.RestBaseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), baseResult)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if baseResult.Code == 0 {
		result := &market.GetMarketTradeResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *MarketClient) GetKlines(paramMap map[string]interface{}) (*market.GetKlinesResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.publicUrlBuilder.Build(constant.V1GetKlinesUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, "")

	result := &market.GetKlinesResponse{}

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

	return result, errors.New(getResp)
}

func (c *MarketClient) GetDeliveryInfo(paramMap map[string]interface{}) (*market.GetDeliveryInfoResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.publicUrlBuilder.Build(constant.V1GetDeliveryInfoUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, "")
	if getErr != nil {
		return nil, getErr
	}
	baseResult := &base.RestBaseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), baseResult)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if baseResult.Code == 0 {
		result := &market.GetDeliveryInfoResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *MarketClient) GetMarketSummary(paramMap map[string]interface{}) (*market.GetMarketSummaryResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.publicUrlBuilder.Build(constant.V1GetMarketSummaryUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, "")
	if getErr != nil {
		return nil, getErr
	}
	baseResult := &base.RestBaseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), baseResult)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if baseResult.Code == 0 {
		result := &market.GetMarketSummaryResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *MarketClient) GetFundingRate(instrumentId string) (*market.GetFundingRateResponse, error) {
	//req := new(model.GetRequest).Init()
	//req.AddParam("instrument_id", instrumentId)
	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = instrumentId

	url := c.publicUrlBuilder.Build(constant.V1GetFundingRateUrl, paramMap)
	getResp, getErr := internal.HttpGet(url, "")

	result := &market.GetFundingRateResponse{}

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
