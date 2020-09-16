package wsclient

import (
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
)

type PrivateWebsocketClient struct {
	WebSocketClientBase
}

func (p *PrivateWebsocketClient) Init(host, token string) *PrivateWebsocketClient {
	p.WebSocketClientBase.Init(host, token)
	return p
}

func (p *PrivateWebsocketClient) SetHandler(
	connectedHandler ConnectedHandler,
	responseHandler ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, responseHandler)
}

func (p *PrivateWebsocketClient) Subscribe(paramMap map[string]interface{}) {
	requestBody, err := p.BuildRequestBody(paramMap)
	if err != nil {
		applogger.Error("Build param body error: %s", err)
		return
	}
	applogger.Info("Send subscribe message: %s", requestBody)
	p.Send(requestBody)
}

func (p *PrivateWebsocketClient) UnSubscribe(paramMap map[string]interface{}) {
	requestBody, err := p.BuildRequestBody(paramMap)
	if err != nil {
		applogger.Error("Build param body error: %s", err)
		return
	}
	applogger.Info("Send unsubscribe message: %s", requestBody)
	p.Send(requestBody)
}
