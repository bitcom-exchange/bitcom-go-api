package blocktrade

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type QueryBlockTradeResponse struct {
	base.RestBaseResponse
	Data []ParadigmBlockTradeQueryVo `json:"data"`
}

type ParadigmBlockTradeQueryVo struct {
	BlockOrderId string `json:"block_order_id" example:"1001"`
	Label        string `json:"label" example:"abc"`
	CreatedAt    int64  `json:"created_at" example:"1585296000000"`
	UpdatedAt    int64  `json:"updated_at" example:"1585296000000"`
	UserId       string `json:"user_id" example:"801"`
	Counterparty string `json:"counterparty" example:"1002"`
	InstrumentId string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Side         string `json:"side" example:"buy"`
	Price        string `json:"price" example:"0.03"`
	Qty          string `json:"qty" example:"1"`
	Fee          string `json:"fee" example:"0.0001"`
	Status       string `json:"status" example:"filled"`
	Role         string `json:"role" example:"taker"`
	BtSource     string `json:"bt_source"`
	OrderId      string `json:"order_id"`
}
