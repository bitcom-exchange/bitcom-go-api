package ws

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type IndexPriceResponse struct {
	base.WsBaseResponse
	Data IndexMessage `json:"data"`
}

type IndexMessage struct {
	IndexName  string `json:"index_name"`
	IndexPrice string `json:"index_price"`
}
