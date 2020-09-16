package account

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetUserDeliveriesResponse struct {
	base.RestBaseResponse
	Data []DeliveryVo `json:"data"`
}

type DeliveryVo struct {
	Type          string `json:"type" example:"delivery"`
	Timestamp     int64  `json:"timestamp" example:"1585296000000"`
	InstrumentId  string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Position      string `json:"position" example:"0.7"`
	Exercise      bool   `json:"exercise"`
	DeliveryPrice string `json:"delivery_price" example:"8000"`
	DeliveryPnl   string `json:"delivery_pnl" example:"0.07"`

	//MarkPrice         string `json:"mark_price" example:"0.0035"`
	//SessionProfitLoss string `json:"session_profit_loss" example:"0.35"`
	//ProfitLoss        string `json:"profit_loss" example:"0.07"`
}
