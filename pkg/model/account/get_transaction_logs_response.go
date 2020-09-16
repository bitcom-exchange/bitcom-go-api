package account

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetTransactionLogsResponse struct {
	base.RestBaseResponse
	Data []TransactionLogVo `json:"data"`
}

type TransactionLogVo struct {
	TransactionTime int64  `json:"transaction_time" example:"1585296000000"`
	InstrumentID    string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	TransactionType string `json:"transaction_type" example:"trade"`
	Direction       string `json:"direction" example:"buy"`
	Qty             string `json:"qty" example:"0.1"`
	Price           string `json:"price" example:"0.035"`
	CashFlow        string `json:"cash_flow" example:"0.003"`
	Funding         string `json:"funding" example:"0.00003"`
	FeePaid         string `json:"fee_paid" example:"0.0001"`
	FeeRate         string `json:"fee_rate" example:"0.0002"`
	Change          string `json:"change" example:"-0.02"`
	Balance         string `json:"balance" example:"98.2"`
	TradeID         string `json:"trade_id" example:"3472"`
	OrderID         string `json:"order_id" example:"8721"`
	Position        string `json:"position" example:"3.1"`

	// Percentage Share: xx% for socialized-fund type
	// autoliquidation for liquidation trade
	Remark string `json:"remark" example:"autoliquidation"`
}
