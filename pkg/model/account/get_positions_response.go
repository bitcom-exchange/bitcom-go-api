package account

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetPositionsResponse struct {
	base.RestBaseResponse
	Paging base.OldPaging `json:"paging"`
	Data   []PositionVo   `json:"data"`
}

type PositionVo struct {
	InstrumentId      string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	ExpirationAt      int64  `json:"expiration_at" example:"1585296000000"`
	Qty               string `json:"qty" example:"5"`
	InitialMargin     string `json:"initial_margin" example:"0.023"`
	MaintenanceMargin string `json:"maintenance_margin" example:"0.0075"`
	AvgPrice          string `json:"avg_price" example:"0.016"`
	//UPL               string `json:"upl" example:"0.0024"`
	//RPL               string `json:"rpl" example:"-0.0002"`
	//PNL               string `json:"pnl" example:"0.0021"`
	SessionAvgPrice string `json:"session_avg_price" example:"0.012"`
	MarkPrice       string `json:"mark_price" example:"0.0401"`
	IndexPrice      string `json:"index_price" example:"7800"`
	QtyBase         string `json:"qty_base,omitempty" example:"0.001"`

	LiqPrice           string `json:"liq_price" example:"6950"`
	SessionFunding     string `json:"session_funding,omitempty" example:"0.000023"`
	PositionPnl        string `json:"position_pnl" example:"0.0021"`
	PositionSessionUpl string `json:"position_session_upl" example:"0.0024"`
	PositionSessionRpl string `json:"position_session_rpl" example:"-0.0002"`

	// NEW
	Category string `json:"category" example:"option"`
	ROI      string `json:"roi" example:"0.0003"`

	OptionDelta string `json:"option_delta,omitempty"`
	OptionGamma string `json:"option_gamma,omitempty"`
	OptionVega  string `json:"option_vega,omitempty"`
	OptionTheta string `json:"option_theta,omitempty"`

	Leverage string `json:"leverage" example:"30"`
}
