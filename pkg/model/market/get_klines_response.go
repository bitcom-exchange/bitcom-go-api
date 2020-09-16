package market

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetKlinesResponse struct {
	base.RestBaseResponse
	Data KlineVo `json:"data"`
}

type KlineVo struct {
	Volume []float64 `json:"volume"`
	//Ticks  []int64   `json:"ticks"`
	Timestamps []int64 `json:"timestamps"`
	//Status string    `json:"status"`
	Open  []float64 `json:"open"`
	Low   []float64 `json:"low"`
	High  []float64 `json:"high"`
	Close []float64 `json:"close"`
}
