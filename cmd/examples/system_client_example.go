package examples

import (
	"github.com/bitcom-exchange/bitcom-go-api/config"
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/restclient"
)

func GetSystemVersionExample() {
	systemClient := new(restclient.SystemClient).Init(config.User1Host)
	resp, err := systemClient.GetSystemVersion()
	if err != nil {
		applogger.Error("Get system version error: %s", err)
	} else {
		applogger.Info("Get system version: %s", resp.Data)
	}
}

func GetSystemTimestampExample() {
	systemClient := new(restclient.SystemClient).Init(config.User1Host)
	resp, err := systemClient.GetServerTimestamp()
	if err != nil {
		applogger.Error("Get system timestamp error: %s", err)
	} else {
		applogger.Info("Get system timestamp: %d", resp.Data)
	}
}
