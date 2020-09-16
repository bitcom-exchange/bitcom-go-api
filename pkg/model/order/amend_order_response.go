package order

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type AmendOrderResponse struct {
	base.RestBaseResponse
	Data OrderActionVo `json:"data"`
}
