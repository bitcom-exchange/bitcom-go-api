package order

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type AmendBatchOrdersResponse struct {
	base.RestBaseResponse
	Data BatchOrderVo `json:"data"`
}
