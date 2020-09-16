package ws

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type MarkPriceResponse struct {
	base.WsBaseResponse
	Data MarkPriceMessage `json:"data"`
}

type MarkPriceMessage struct {
	InstrumentId string `json:"instrument_id"`
	MarkPrice    string `json:"mark_price"`
	Sigma        string `json:"sigma"`
	Delta        string `json:"delta"`
	Vega         string `json:"vega"`
	Theta        string `json:"theta"`
	Gama         string `json:"gama"`
}
