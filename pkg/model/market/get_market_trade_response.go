package market

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetMarketTradeResponse struct {
	base.RestBaseResponse
	PageInfo base.NewPaging  `json:"page_info"`
	Data     []MarketTradeVo `json:"data"`
}

type MarketTradeVo struct {
	CreatedAt       int64  `json:"created_at" example:"1585299600000"`
	TradeId         int64  `json:"trade_id" example:"3743"`
	InstrumentId    string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Price           string `json:"price" example:"0.034"`
	Qty             string `json:"qty" example:"1"`
	Side            string `json:"side" example:"buy"`
	Sigma           string `json:"sigma" example:"0.002"`
	IndexPrice      string `json:"index_price" example:"8000"`
	UnderlyingPrice string `json:"underlying_price" example:"7900"`
	IsBlockTrade    bool   `json:"is_block_trade" example:"false"`
}
