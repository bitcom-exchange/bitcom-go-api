package main

import "github.com/bitcom-exchange/bitcom-go-api/cmd/examples"

func main() {
	// system client
	//examples.GetSystemVersionExample()
	//examples.GetSystemTimestampExample()

	// market client
	//examples.GetIndexExample()
	//examples.GetInstrumentsExample()
	//examples.GetTickerExample()
	//examples.GetOrderBookExample()
	//examples.GetMarketTradeExample()
	//examples.GetKlinesExample()
	//examples.GetDeliveryInfoExample()
	//examples.GetMarketSummaryExample()
	//examples.GetFundingRateExample()

	// account client
	//examples.GetAccountsExample()
	//examples.GetPositionsExample()
	//examples.GetTransactionLogsExample()
	//examples.GetUserDeliveriesExample()
	//examples.GetUserSettlementsExample()
	//examples.ConfigCodExample()
	//examples.GetCodConfigExample()
	//examples.GetMmpStateExample()
	//examples.UpdateMmpConfigExample()
	examples.ResetMmpStateExample()

	// order client
	//examples.PlaceNewOrderExample()
	//examples.PlaceNewBatchOrderExample()
	//examples.CancelOrderExample()
	//examples.AmendOrderExample()
	//examples.AmendBatchOrdersExample()
	//examples.ClosePositionsExample()
	//examples.GetOpenOrdersExample()
	//examples.GetOrdersExample()
	//examples.GetStopOrdersExample()
	//examples.GetUserTradesExample()
	//examples.GetEstMarginsExample()

	// Block Trade
	//examples.NewParadigmBlockOrdersExample()
	//examples.QueryParadigmBlockOrdersExample()
	//examples.QueryParadigmBlockOrdersByPlatformExample()

	// WebSocket
	//examples.PublicSubscribeExample()
	//examples.PrivateSubscribeExample()
}
