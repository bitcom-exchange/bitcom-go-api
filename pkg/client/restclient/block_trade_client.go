package restclient

import (
	"encoding/json"
	"errors"
	"github.com/bitcom-exchange/bitcom-go-api/constant"
	"github.com/bitcom-exchange/bitcom-go-api/internal"
	"github.com/bitcom-exchange/bitcom-go-api/internal/requestbuilder"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/blocktrade"
)

type BlockTradeClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

func (c *BlockTradeClient) Init(host string, accessKey string, secretKey string) *BlockTradeClient {
	c.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(host, accessKey, secretKey)
	return c
}

func (c *BlockTradeClient) NewParadigmBlockOrders(paramMap map[string]interface{}) (*blocktrade.NewBlockTradeResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url, postBody, err := c.privateUrlBuilder.BuildWithBody(constant.V1NewParadigmBlockOrdersUrl, paramMap)
	if err != nil {
		return nil, err
	}

	postResp, postErr := internal.HttpPost(url, postBody, c.privateUrlBuilder.GetAccessKey())
	if postErr != nil {
		return nil, postErr
	}
	result := &blocktrade.NewBlockTradeResponse{}

	jsonErr := json.Unmarshal([]byte(postResp), result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 0 {
		return result, nil
	}

	return nil, errors.New(postResp)
}

func (c *BlockTradeClient) QueryParadigmBlockOrders(paramMap map[string]interface{}) (*blocktrade.QueryBlockTradeResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1QueryParadigmBlockOrdersUrl, paramMap)

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
		result := &blocktrade.QueryBlockTradeResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}

func (c *BlockTradeClient) QueryParadigmBlockOrdersByPlatform(paramMap map[string]interface{}) (*blocktrade.QueryBlockTradeByPlatformResponse, error) {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	url := c.privateUrlBuilder.BuildWithURL(constant.V1QueryBlockOrdersByPlatformUrl, paramMap)

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
		result := &blocktrade.QueryBlockTradeByPlatformResponse{}
		jsonErr := json.Unmarshal([]byte(getResp), result)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return result, nil
	}

	return nil, errors.New(getResp)
}
