package account

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetCodResponse struct {
	base.RestBaseResponse
	Data CodVO `json:"data"`
}

type PostCodResponse struct {
	base.RestBaseResponse
	Data interface{} `json:"data"`
}

type CodVO struct {
	Cod bool `json:"cod"`
}
