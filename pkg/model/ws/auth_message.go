package ws

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetWsAuthTokenResponse struct {
	base.RestBaseResponse
	Data WsToken `json:"data"`
}

type WsToken struct {
	Token string `json:"token"`
}
