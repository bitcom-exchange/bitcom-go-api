package order

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type ClosePositionsResponse struct {
	base.RestBaseResponse
	Data OrderActionVo `json:"data"`
}
