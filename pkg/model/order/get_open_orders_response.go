package order

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetOpenOrdersResponse struct {
	base.RestBaseResponse
	Data []OrderVo `json:"data"`
}
