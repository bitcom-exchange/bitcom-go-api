package requestbuilder

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

type PrivateUrlBuilder struct {
	host      string
	accessKey string
	secretKey string
	tKey      string

	signer *Signer
}

func (p *PrivateUrlBuilder) Init(host string, accessKey string, secretKey string) *PrivateUrlBuilder {
	p.tKey = "timestamp"
	p.accessKey = accessKey
	p.secretKey = secretKey
	p.host = host
	p.signer = new(Signer).Init(secretKey)
	return p
}

func (p *PrivateUrlBuilder) GetAccessKey() string {
	return p.accessKey
}

func (p *PrivateUrlBuilder) GetSecretKey() string {
	return p.secretKey
}

func (p *PrivateUrlBuilder) BuildWithURL(path string, paramMap map[string]interface{}) string {
	utc := time.Now().UTC()
	timestamp := utc.UnixNano() / 1e6

	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	paramMap["timestamp"] = timestamp

	var paramList []string
	for k, v := range paramMap {
		paramList = append(paramList, fmt.Sprintf("%v=%v", k, v))
	}
	requestUrl := path + "?" + strings.Join(paramList, "&")

	//paramToSign := p.buildParamToSign(paramMap)
	signature := p.signer.Sign(path, paramMap)

	requestUrl = fmt.Sprintf("%s%s&signature=%s", p.host, requestUrl, url.QueryEscape(signature))

	return requestUrl
}

func (p PrivateUrlBuilder) BuildWithBody(path string, paramMap map[string]interface{}) (string, string, error) {
	utc := time.Now().UTC()
	timestamp := utc.UnixNano() / 1e6

	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	paramMap["timestamp"] = timestamp
	//paramToSign := p.buildParamToSign(paramMap)
	signature := p.signer.Sign(path, paramMap)
	paramMap["signature"] = signature

	var jsonBody string

	if len(paramMap) > 0 {
		jsonBuf, err := json.Marshal(paramMap)
		if err != nil {
			return "", "", err
		}

		jsonBody = string(jsonBuf)
	}

	requestUrl := fmt.Sprintf("%s%s", p.host, path)
	return requestUrl, jsonBody, nil
}

func (p *PrivateUrlBuilder) buildParamToSign(paramMap map[string]interface{}) string {
	// To store the keys in slice in sorted order
	var keys []string
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder

	// To perform the opertion you want
	for _, k := range keys {
		sb.WriteString("&")
		sb.WriteString(k)
		sb.WriteString("=")
		strValue := fmt.Sprintf("%v", paramMap[k])
		sb.WriteString(strValue)
	}

	return sb.String()
}
