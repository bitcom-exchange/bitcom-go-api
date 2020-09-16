package ws

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type MarketTradeResponse struct {
	base.WsBaseResponse
	Data []TradeMessage `json:"data"`
}
