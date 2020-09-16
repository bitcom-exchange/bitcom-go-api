package account

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetAccountResponse struct {
	base.RestBaseResponse
	Data AccountVo `json:"data"`
}

type AccountVo struct {
	UserID            string `json:"user_id" example:"9001"`
	Currency          string `json:"currency" example:"BTC"`
	CashBalance       string `json:"cash_balance" example:"10.00000000"`
	AvailableBalance  string `json:"available_balance" example:"100.00000000"`
	MarginBalance     string `json:"margin_balance" example:"90.00000000"`
	InitialMargin     string `json:"initial_margin" example:"10.00000000"`
	MaintenanceMargin string `json:"maintenance_margin" example:"70.00000000"`
	Equity            string `json:"equity" example:"105.00000000"`
	PNL               string `json:"pnl" example:"5.00000000"`
	TotalDelta        string `json:"total_delta" example:"3.12000000"`

	// NEW
	AccountID string `json:"account_id"`
	Mode      string `json:"mode"`
	//AvailableWithdrawal string `json:"available_withdrawal"`
	BlockTradePermission bool `json:"block_trade_permission"`

	SessionUpl string `json:"session_upl"`
	SessionRpl string `json:"session_rpl"`

	OptionValue      string `json:"option_value"`
	OptionPnl        string `json:"option_pnl"`
	OptionSessionRpl string `json:"option_session_rpl"`
	OptionSessionUpl string `json:"option_session_upl"`
	OptionDelta      string `json:"option_delta"`
	OptionGamma      string `json:"option_gamma"`
	OptionVega       string `json:"option_vega"`
	OptionTheta      string `json:"option_theta"`

	FuturePnl            string `json:"future_pnl"`
	FutureSessionRpl     string `json:"future_session_rpl"`
	FutureSessionUpl     string `json:"future_session_upl"`
	FutureSessionFunding string `json:"future_session_funding"`
	FutureDelta          string `json:"future_delta"`

	CreatedAt int64 `json:"created_at"`
}
