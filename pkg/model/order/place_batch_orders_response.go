package order

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type NewBatchOrdersResponse struct {
	base.RestBaseResponse
	Data BatchOrderVo `json:"data"`
}

type BatchOrderDetailVo struct {
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
	IsLiquidation bool   `json:"is_liquidation" example:"false"`
	AutoPrice     string `json:"auto_price" example:"7000"`
	AutoPriceType string `json:"auto_price_type" example:"usd"`
	TakerFeeRate  string `json:"taker_fee_rate" example:"0.00005"`
	MakerFeeRate  string `json:"maker_fee_rate" example:"0.00002"`
	Label         string `json:"label" example:"strategy-A"`
	//StopPrice     string `json:"stop_price"`
	ReduceOnly     bool   `json:"reduce_only"`
	PostOnly       bool   `json:"post_only"`
	RejectPostOnly bool   `json:"reject_post_only"`
	ErrorCode      int    `json:"error_code"`
	ErrorMsg       string `json:"error_msg"`

	Mmp    bool   `json:"mmp"`
	Source string `json:"source"`
	Hidden bool   `json:"hidden"`
}

type BatchOrderVo struct {
	Orders []BatchOrderDetailVo `json:"Orders"`
}
