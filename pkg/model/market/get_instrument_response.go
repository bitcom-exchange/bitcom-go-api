package market

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetInstrumentsResponse struct {
	base.RestBaseResponse
	Data []InstrumentVo `json:"data"`
}

type InstrumentVo struct {
	InstrumentId  string `json:"instrument_id" example:"BTC-27MAR20-9000-C"`
	CreatedAt     int64  `json:"created_at" example:"1582790400000"`
	UpdatedAt     int64  `json:"updated_at" example:"1582790400000"`
	BaseCurrency  string `json:"base_currency" example:"BTC"`
	QuoteCurrency string `json:"quote_currency" example:"USD"`
	StrikePrice   string `json:"strike_price" example:"9000"`
	ExpirationAt  int64  `json:"expiration_at" example:"1585296000000"`
	OptionType    string `json:"option_type" example:"call"`

	//BaseMinQty     string `json:"base_min_qty" example:"1"`
	//BaseMaxQty     string `json:"base_max_qty" example:"9999"`
	//BasePrecision  int    `json:"base_precision" example:"8"`
	//QuotePrecision int    `json:"quote_precision" example:"8"`
	//QtyPrecision   int    `json:"qty_precision" example:"8"`
	//TickSize       string `json:"tick_size" example:"4"`

	Category             string `json:"category"`
	MinPrice             string `json:"min_price" example:"0.0005"`
	MaxPrice             string `json:"max_price" example:"10"`
	PriceStep            string `json:"price_step" example:"0.0005"`
	MinSize              string `json:"min_size" example:"0.1"`
	SizeStep             string `json:"size_step" example:"0.1"`
	DeliveryFeeRate      string `json:"delivery_fee_rate" example:"0.0002"`
	ContractSize         string `json:"contract_size"`
	ContractSizeCurrency string `json:"contract_size_currency"`
	Active               bool   `json:"active"`
}
