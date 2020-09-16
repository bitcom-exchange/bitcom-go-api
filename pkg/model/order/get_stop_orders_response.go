package order

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetStopOrdersResponse struct {
	base.RestBaseResponse
	Data []StopOrderVo `json:"data"`
}

type StopOrderVo struct {
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	Status       string `json:"status"`
	StopPrice    string `json:"stop_price"`
	TriggerType  string `json:"trigger_type"`
	RejectReason string `json:"reject_reason"`

	// fields of order
	//OrderId      string `json:"order_id"`
	StopOrderId  string `json:"stop_order_id"`
	InstrumentId string `json:"instrument_id"`
	UserId       string `json:"user_id"`
	Qty          string `json:"qty"`
	Price        string `json:"price"`
	Side         string `json:"side"`
	OrderType    string `json:"order_type"`
	TimeInForce  string `json:"time_in_force"`
}
