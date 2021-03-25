package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string, accessKey string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	if accessKey != "" {
		req.Header.Add("X-Bit-Access-Key", accessKey)
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Http get request failed.\n response body is %v ", string(result))
	}
	return string(result), nil
}

func HttpPost(url string, body string, accessKey string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	if accessKey != "" {
		req.Header.Add("X-Bit-Access-Key", accessKey)
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Http post request failed.\n response body is %v ", string(result))
	}
	return string(result), nil
}
