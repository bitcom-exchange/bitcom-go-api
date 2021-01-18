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

type GetCancelOnlyStatusResponse struct {
	base.RestBaseResponse
	Data CancelOnlyStatusVo `json:"data"`
}

type CancelOnlyStatusVo struct {
	Status   int64 `json:"status"`
	RemainMs int64 `json:"remain_ms"`
}
