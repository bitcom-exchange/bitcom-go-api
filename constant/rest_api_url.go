package constant

const (
	// System
	V1GetServerTimestampUrl  = "/v1/system/time"
	V1GetSystemVersionUrl    = "/v1/system/version"
	V1GetCancelOnlyStatusUrl = "/v1/system/cancel_only_status"

	// Market
	V1GetIndexUrl              = "/v1/index"
	V1GetInstrumentsUrl        = "/v1/instruments"
	V1GetTickersUrl            = "/v1/tickers"
	V1GetOrderbooksUrl         = "/v1/orderbooks"
	V1GetMarketTradesUrl       = "/v1/market/trades"
	V1GetKlinesUrl             = "/v1/klines"
	V1GetDeliveryInfoUrl       = "/v1/delivery_info"
	V1GetMarketSummaryUrl      = "/v1/market/summary"
	V1GetFundingRateUrl        = "/v1/funding_rate"
	V1GetFundingRateHistoryUrl = "/v1/funding_rate_history"
	V1GetTotalVolumeUrl        = "/v1/total_volumes"

	// Account
	V1GetAccountsUrl        = "/v1/accounts"
	V1GetPositionsUrl       = "/v1/positions"
	V1GetTransactionLogsUrl = "/v1/transactions"
	V1GetUserDeliveriesUrl  = "/v1/user/deliveries"
	V1GetUserSettlementsUrl = "/v1/user/settlements"
	V1ConfigCodUrl          = "/v1/account_configs/cod"
	V1GetCodConfigUrl       = "/v1/account_configs/cod"
	V1GetMmpStateUrl        = "/v1/mmp_state"
	V1UpdateMmpConfigUrl    = "/v1/update_mmp_config"
	V1ResetMmpStateUrl      = "/v1/reset_mmp"

	// Order
	V1NewOrderUrl       = "/v1/orders"
	V1CancelOrdersUrl   = "/v1/cancel_orders"
	V1AmendOrdersUrl    = "/v1/amend_orders"
	V1ClosePositionsUrl = "/v1/close_positions"
	V1GetOpenOrdersUrl  = "/v1/open_orders"
	V1GetOrdersUrl      = "/v1/orders"
	V1GetStopOrdersUrl  = "/v1/stop_orders"
	V1GetUserTradesUrl  = "/v1/user/trades"
	V1GetEstMarginsUrl  = "/v1/margins"

	// Batch Orders
	V1NewBatchOrdersUrl   = "/v1/batchorders"
	V1AmendBatchOrdersUrl = "/v1/amend_batchorders"

	// Block Trade
	V1NewParadigmBlockOrdersUrl     = "/v1/blocktrades"
	V1QueryParadigmBlockOrdersUrl   = "/v1/blocktrades"
	V1QueryBlockOrdersByPlatformUrl = "/v1/platform_blocktrades"

	// Websocket Auth
	V1WsAuthToken = "/v1/ws/auth"
)
