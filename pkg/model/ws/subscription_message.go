package ws

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type SubscriptionSuccessResponse struct {
	base.WsBaseResponse
	Data SubscriptionSuccessMessage `json:"data"`
}

type SubscriptionFailResponse struct {
	base.WsBaseResponse
	Data SubscriptionFailMessage `json:"data"`
}

type SubscriptionSuccessMessage struct {
	Code         int64    `json:"code"`
	Subscription []string `json:"subscription"`
}

type SubscriptionFailMessage struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
