package market

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetTickerResponse struct {
	base.RestBaseResponse
	Data TickerVo `json:"data"`
}

type TickerVo struct {
	Time         int64  `json:"time"`
	InstrumentId string `json:"instrument_id"`

	BestBid    string `json:"best_bid"`
	BestAsk    string `json:"best_ask"`
	BestBidQty string `json:"best_bid_qty"`
	BestAskQty string `json:"best_ask_qty"`
	AskSigma   string `json:"ask_sigma"`
	BidSigma   string `json:"bid_sigma"`

	LastPrice      string `json:"last_price"`
	LastQty        string `json:"last_qty"`
	Open24H        string `json:"open24h"`
	High24H        string `json:"high24h"`
	Low24H         string `json:"low24h"`
	PriceChange24H string `json:"price_change24h"`
	Volume24H      string `json:"volume24h"`
	OpenInterest   string `json:"open_interest"`

	UnderlyingName  string `json:"underlying_name,omitempty"`
	UnderlyingPrice string `json:"underlying_price,omitempty"`
	MarkPrice       string `json:"mark_price"`
	Sigma           string `json:"sigma,omitempty"`
	Delta           string `json:"delta,omitempty"`
	Vega            string `json:"vega,omitempty"`
	Theta           string `json:"theta,omitempty"`
	Gamma           string `json:"gamma,omitempty"`

	MinSellPrice string `json:"min_sell"`
	MaxBuyPrice  string `json:"max_buy"`
}
