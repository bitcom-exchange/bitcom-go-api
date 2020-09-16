package market

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetMarketSummaryResponse struct {
	base.RestBaseResponse
	Data []MarketSummaryVo `json:"data"`
}

type MarketSummaryVo struct {
	InstrumentId string `json:"instrument_id"`
	Timestamp    int64  `json:"timestamp"`

	BestBid    string `json:"best_bid"`
	BestAsk    string `json:"best_ask"`
	BestBidQty string `json:"best_bid_qty"`
	BestAskQty string `json:"best_ask_qty"`

	LastPrice    string `json:"last_price"`
	LastQty      string `json:"last_qty"`
	Open24H      string `json:"open24h"`
	High24H      string `json:"high24h"`
	Low24H       string `json:"low24h"`
	Volume24H    string `json:"volume24h"`
	OpenInterest string `json:"open_interest"`

	MarkPrice string `json:"mark_price"`

	MaxBuy  string `json:"max_buy"`
	MinSell string `json:"min_sell"`

	Delta string `json:"delta"`
	Gamma string `json:"gamma"`
	Vega  string `json:"vega"`
	Theta string `json:"theta"`
	//Sigma string `json:"sigma"`
}
