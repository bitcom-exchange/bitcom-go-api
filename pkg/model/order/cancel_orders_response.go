package order

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type NewCancelOrdersResponse struct {
	base.RestBaseResponse
	Data OrderCancelVo `json:"data"`
}

type OrderCancelVo struct {
	NumCancelled int64 `json:"num_cancelled"`
}
