package ws

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type OrderBookResponse struct {
	base.WsBaseResponse
	Data OrderBookMessage `json:"data"`
}

type OrderBookMessage struct {
	InstrumentId string      `json:"instrument_id"`
	Sequence     int64       `json:"sequence"`
	Bids         [][2]string `json:"bids"` // [[<price>, <qty>]]
	Asks         [][2]string `json:"asks"`
}
