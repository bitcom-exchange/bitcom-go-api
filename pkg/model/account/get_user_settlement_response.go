package account

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetUserSettlementsResponse struct {
	base.RestBaseResponse
	Data []SettlementVo `json:"data"`
}

type SettlementVo struct {
	Type           string `json:"type" example:"delivery"`
	Timestamp      int64  `json:"timestamp" example:"1585296000000"`
	InstrumentId   string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Position       string `json:"position" example:"0.7"`
	Direction      string `json:"direction"`
	SessionUpl     string `json:"session_upl" example:"0.0002"`
	SessionRpl     string `json:"session_rpl" example:"-0.0006"`
	SessionFunding string `json:"session_funding" example:"0.0000002"`
	//MarkPrice      string `json:"mark_price" example:"0.0035"`
	//SessionPnl     string `json:"session_pnl" example:"0.0003"`
}
