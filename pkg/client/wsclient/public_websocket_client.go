package wsclient

import (
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
)

type PublicWebsocketClient struct {
	WebSocketClientBase
}

func (p *PublicWebsocketClient) Init(host string) *PublicWebsocketClient {
	p.WebSocketClientBase.Init(host, "")
	return p
}

func (p *PublicWebsocketClient) SetHandler(
	connectedHandler ConnectedHandler,
	responseHandler ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, responseHandler)
}

func (p *PublicWebsocketClient) Subscribe(paramMap map[string]interface{}) {
	requestBody, err := p.BuildRequestBody(paramMap)
	if err != nil {
		applogger.Error("Build param body error: %s", err)
		return
	}
	applogger.Info("Send subscribe message: %s", requestBody)
	p.Send(requestBody)
}

func (p *PublicWebsocketClient) UnSubscribe(paramMap map[string]interface{}) {
	requestBody, err := p.BuildRequestBody(paramMap)
	if err != nil {
		applogger.Error("Build param body error: %s", err)
		return
	}
	applogger.Info("Send unsubscribe message: %s", requestBody)
	p.Send(requestBody)
}
