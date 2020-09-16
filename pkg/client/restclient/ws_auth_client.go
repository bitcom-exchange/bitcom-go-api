package restclient

import (
	"encoding/json"
	"errors"
	"github.com/bitcom-exchange/bitcom-go-api/constant"
	"github.com/bitcom-exchange/bitcom-go-api/internal"
	"github.com/bitcom-exchange/bitcom-go-api/internal/requestbuilder"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/ws"
)

type WsAuthClient struct {
	privateBuilder *requestbuilder.PrivateUrlBuilder
}

func (c *WsAuthClient) Init(host string, accessKey string, secretKey string) *WsAuthClient {
	c.privateBuilder = new(requestbuilder.PrivateUrlBuilder).Init(host, accessKey, secretKey)
	return c
}

func (c *WsAuthClient) GetWsAuthToken() (*ws.GetWsAuthTokenResponse, error) {
	url := c.privateBuilder.BuildWithURL(constant.V1WsAuthToken, nil)
	getResp, getErr := internal.HttpGet(url, c.privateBuilder.GetAccessKey())

	result := &ws.GetWsAuthTokenResponse{}

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
