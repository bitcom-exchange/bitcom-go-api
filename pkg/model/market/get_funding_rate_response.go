package market

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetFundingRateResponse struct {
	base.RestBaseResponse
	Data FundingRateVo `json:"data"`
}

type FundingRateVo struct {
	InstrumentId string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	Time         int64  `json:"time" example:"1585296000000"`
	FundingRate  string `json:"funding_rate" example:"0.0000023"`
	IndexPrice   string `json:"index_price" example:"8000"`
	MarkPrice    string `json:"mark_price" example:"7993"`
}
