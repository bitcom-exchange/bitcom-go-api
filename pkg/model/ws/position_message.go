package ws

import (
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/account"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"
)

type PositionResponse struct {
	base.WsBaseResponse
	Data []PositionVo `json:"data"`
}

type PositionVo account.PositionVo
