package internal

import (
	"github.com/bitcom-exchange/bitcom-go-api/logging/perflogger"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string, accessKey string) (string, error) {

	logger := perflogger.GetInstance()
	logger.Start()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	if accessKey != "" {
		req.Header.Add("X-Bit-Access-Key", accessKey)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("GET", url)

	return string(result), err
}

func HttpPost(url string, body string, accessKey string) (string, error) {
	logger := perflogger.GetInstance()
	logger.Start()

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	if accessKey != "" {
		req.Header.Add("X-Bit-Access-Key", accessKey)
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("POST", url)

	return string(result), err
}
