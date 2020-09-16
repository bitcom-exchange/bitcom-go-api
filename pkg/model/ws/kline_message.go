package ws

import (
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"
)

type KlineResponse struct {
	base.WsBaseResponse
	Data KlineMessage `json:"data"`
}

type KlineMessage struct {
	InstrumentId string `json:"instrument_id"`
	Tick         int64  `json:"tick"`
	Open         string `json:"open"`
	Close        string `json:"close"`
	High         string `json:"high"`
	Low          string `json:"low"`
	Volume       string `json:"volume"`
}
