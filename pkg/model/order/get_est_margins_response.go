package order

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetEstMarginsResponse struct {
	base.RestBaseResponse
	Data EstMarginVo `json:"data"`
}

type EstMarginVo struct {
	BuyMargin  string `json:"buy_margin"`
	SellMargin string `json:"sell_margin"`
	MinSell    string `json:"min_sell"`
	MaxBuy     string `json:"max_buy"`
}
