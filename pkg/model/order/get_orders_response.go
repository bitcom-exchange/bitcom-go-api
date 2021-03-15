package order

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetOrdersResponse struct {
	base.RestBaseResponse
	PageInfo base.NewPaging `json:"page_info"`
	Data     []OrderVo      `json:"data"`
}

type OrderVo struct {
	OrderId       string `json:"order_id" example:"1001"`
	CreatedAt     int64  `json:"created_at" example:"1585296000000"`
	UpdatedAt     int64  `json:"updated_at" example:"1585296000000"`
	UserId        string `json:"user_id" example:"801"`
	InstrumentId  string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	OrderType     string `json:"order_type" example:"limit"`
	Side          string `json:"side" example:"buy"`
	Price         string `json:"price" example:"0.03"`
	Qty           string `json:"qty" example:"1"`
	TimeInForce   string `json:"time_in_force" example:"gtc"`
	AvgPrice      string `json:"avg_price" example:"0.029"`
	FilledQty     string `json:"filled_qty" example:"1"`
	Status        string `json:"status" example:"filled"`
	Fee           string `json:"fee" example:"0.00002"`
	IsLiquidation bool   `json:"is_liquidation" example:"false"`
	AutoPrice     string `json:"auto_price" example:"7000"`
	AutoPriceType string `json:"auto_price_type" example:"usd"`
	PNL           string `json:"pnl" example:"0.031"`
	CashFlow      string `json:"cash_flow" example:"0.027"`

	InitialMargin  string `json:"initial_margin" example:"0.04"`
	TakerFeeRate   string `json:"taker_fee_rate" example:"0.00005"`
	MakerFeeRate   string `json:"maker_fee_rate" example:"0.00002"`
	Label          string `json:"label" example:"strategy-A"`
	StopOrderId    string `json:"stop_order_id" example:"stop-x3gjsdhf3232"`
	StopPrice      string `json:"stop_price" example:"9800"`
	ReduceOnly     bool   `json:"reduce_only"`
	PostOnly       bool   `json:"post_only"`
	RejectPostOnly bool   `json:"reject_post_only"`

	Mmp    bool   `json:"mmp"`
	Source string `json:"source"`
	Hidden bool   `json:"hidden"`
}
