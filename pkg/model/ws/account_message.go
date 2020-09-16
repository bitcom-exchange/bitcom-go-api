package ws

import (
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/account"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"
)

type AccountResponse struct {
	base.WsBaseResponse
	Data AccountVo `json:"data"`
}

type AccountVo account.AccountVo
