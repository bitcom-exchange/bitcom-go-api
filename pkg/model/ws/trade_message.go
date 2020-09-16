package ws

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type TradeResponse struct {
	base.WsBaseResponse
	Data []TradeMessage `json:"data"`
}

type TradeMessage struct {
	InstrumentId string `json:"instrument_id"`
	TradeId      string `json:"trade_id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	Side         string `json:"side"`
	Sigma        string `json:"sigma"`
	OptionType   string `json:"option_type"`
	IsBlockTrade bool   `json:"is_block_trade"`
	CreatedAt    int64  `json:"created_at"`
}
