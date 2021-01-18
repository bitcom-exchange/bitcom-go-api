package constant

const (
	// Currency
	BTC = "BTC"
	ETH = "ETH"
	BCH = "BCH"

	// Instrument
	BTCPerpetual = "BTC-PERPETUAL"

	// Account mode
	RegularAccountMode         = "regular"
	PortfolioMarginAccountMode = "portfolio_margin"

	// Instrument Category
	Option = "option"
	Future = "future"

	// Option type
	CallOption = "call"
	PutOption  = "put"

	// Order side
	BuySide  = "buy"
	SellSide = "sell"

	// Order type
	LimitOrder      = "limit"
	MarketOrder     = "market"
	StopLimitOrder  = "stop-limit"
	StopMarketOrder = "stop-market"

	// Order status
	OrderPendingStatus   = "pending"
	OrderOpenStatus      = "open"
	OrderFilledStatus    = "filled"
	OrderCancelledStatus = "cancelled"

	// StopOrder status
	StopOrderOpenStatus      = "open"
	StopOrderTriggeredStatus = "triggered"
	StopOrderCancelledStatus = "cancelled"

	// Order time in force
	OrderTimeInForceGTC = "gtc"
	OrderTimeInForceFOK = "fok"
	OrderTimeInForceIOC = "ioc"

	// Order auto price type
	USD   = "usd"
	Implv = "implv"

	// Block Trade Source
	ParadigmBlockTradeSource = "paradigm"
	AccxBlockTradeSource     = "accx"

	// Market Role
	MarketMaker = "maker"
	MarketTaker = "taker"
)
