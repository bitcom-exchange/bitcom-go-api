package market

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetOrderBookResponse struct {
	base.RestBaseResponse
	Data OrderBookVo `json:"data"`
}

type OrderBookVo struct {
	InstrumentId string     `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Timestamp    int64      `json:"timestamp" example:"1585299600000"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}
