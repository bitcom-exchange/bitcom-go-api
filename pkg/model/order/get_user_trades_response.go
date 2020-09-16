package order

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetUserTradesResponse struct {
	base.RestBaseResponse
	Data []TradeVo `json:"data"`
}

type TradeVo struct {
	TradeId         string `json:"trade_id" example:"3826"`
	OrderId         string `json:"order_id" example:"1001"`
	InstrumentId    string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Qty             string `json:"qty" example:"1"`
	Price           string `json:"price" example:"0.2275"`
	Sigma           string `json:"sigma" example:"0.0024"`
	UnderlyingPrice string `json:"underlying_price" example:"6750"`
	IndexPrice      string `json:"index_price" example:"6800"`
	UsdPrice        string `json:"usd_price" example:"1664"`
	Fee             string `json:"fee" example:"0.003"`
	FeeRate         string `json:"fee_rate" example:"0.0003"`
	Side            string `json:"side" example:"buy"`
	CreatedAt       int64  `json:"created_at" example:"1585296000000"`
	IsTaker         bool   `json:"is_taker" example:"true"`
	OrderType       string `json:"order_type" example:"limit"`
	IsBlockTrade    bool   `json:"is_block_trade" example:"false"`
	Label           string `json:"label" example:"hedge"`
}
