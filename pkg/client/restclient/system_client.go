package restclient

import (
	"encoding/json"
	"errors"

	"github.com/bitcom-exchange/bitcom-go-api/constant"
	"github.com/bitcom-exchange/bitcom-go-api/internal"
	"github.com/bitcom-exchange/bitcom-go-api/internal/requestbuilder"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/system"
)

type SystemClient struct {
	publicUrlBuilder *requestbuilder.PublicUrlBuilder
}

func (c *SystemClient) Init(host string) *SystemClient {
	c.publicUrlBuilder = new(requestbuilder.PublicUrlBuilder).Init(host)
	return c
}

func (c *SystemClient) GetSystemVersion() (*system.GetSystemVersionResponse, error) {
	url := c.publicUrlBuilder.Build(constant.V1GetSystemVersionUrl, nil)
	getResp, getErr := internal.HttpGet(url, "")

	result := &system.GetSystemVersionResponse{}

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

func (c *SystemClient) GetServerTimestamp() (*system.GetSystemTimestampResponse, error) {
	url := c.publicUrlBuilder.Build(constant.V1GetServerTimestampUrl, nil)
	getResp, getErr := internal.HttpGet(url, "")
	result := &system.GetSystemTimestampResponse{}

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

func (c *SystemClient) GetCancelOnlyStatus() (*system.GetCancelOnlyStatusResponse, error) {
	url := c.publicUrlBuilder.Build(constant.V1GetCancelOnlyStatusUrl, nil)
	getResp, getErr := internal.HttpGet(url, "")
	result := &system.GetCancelOnlyStatusResponse{}

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
