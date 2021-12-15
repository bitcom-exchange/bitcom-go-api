package market

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetCurrenciesResponse struct {
	base.RestBaseResponse
	Data CurrenciesVo `json:"data"`
}

type CurrenciesVo struct {
	Currencies []string `json:"currencies"`
}
