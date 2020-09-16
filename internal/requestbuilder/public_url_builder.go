package requestbuilder

import (
	"fmt"
	"strings"
)

type PublicUrlBuilder struct {
	host string
}

func (p *PublicUrlBuilder) Init(host string) *PublicUrlBuilder {
	p.host = host
	return p
}

func (p *PublicUrlBuilder) Build(path string, paramMap map[string]interface{}) string {
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	var paramList []string
	for k, v := range paramMap {
		paramList = append(paramList, fmt.Sprintf("%v=%v", k, v))
	}
	path = path + "?" + strings.Join(paramList, "&")
	requestUrl := fmt.Sprintf("%s%s", p.host, path)
	return requestUrl
}
