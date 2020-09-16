package system

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetSystemTimestampResponse struct {
	base.RestBaseResponse
	Data int64 `json:"data"`
}

type GetSystemVersionResponse struct {
	base.RestBaseResponse
	Data string `json:"data"`
}
