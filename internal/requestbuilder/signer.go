package requestbuilder

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Signer struct {
	sync.Mutex
	hash hash.Hash
}

func (s *Signer) Init(key string) *Signer {
	s.hash = hmac.New(sha256.New, []byte(key))
	return s
}

func (s *Signer) Sign(path string, paramMap map[string]interface{}) string {
	if path == "" || paramMap == nil {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(path)
	sb.WriteString("&")
	sb.WriteString(s.encodeMap(paramMap))

	strToSign := sb.String()

	return s.sign(strToSign)
}

func (s *Signer) encodeMapList(itemList []map[string]interface{}) string {
	var strList []string
	for _, v := range itemList {
		objVal := s.encodeMap(v)
		objStr := fmt.Sprintf("%v", objVal)
		strList = append(strList, objStr)
	}

	outputStr := strings.Join(strList, "&")
	outputStr = "[" + outputStr + "]"
	return outputStr
}

func (s *Signer) encodeMap(paramMap map[string]interface{}) string {
	var keys []string
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	resultMap := map[string]interface{}{}

	for _, k := range keys {
		value := paramMap[k]
		switch v := value.(type) {
		case []map[string]interface{}:
			listStr := s.encodeMapList(v)
			resultMap[k] = listStr
		case map[string]interface{}:
			mapStr := s.encodeMap(v)
			resultMap[k] = mapStr
		case bool:
			boolStr := strconv.FormatBool(v)
			resultMap[k] = boolStr
		default:
			generalStr := fmt.Sprintf("%v", v)
			resultMap[k] = generalStr
		}
	}

	return s.buildParamToSign(resultMap)
}

func (s *Signer) buildParamToSign(paramMap map[string]interface{}) string {
	var keys []string
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var resultStrList []string

	for _, k := range keys {
		var sb strings.Builder
		sb.WriteString(k)
		sb.WriteString("=")
		strValue := fmt.Sprintf("%v", paramMap[k])
		sb.WriteString(strValue)
		resultStrList = append(resultStrList, sb.String())
	}

	outputStr := strings.Join(resultStrList, "&")
	return outputStr
}

func (s *Signer) sign(strToSign string) string {
	s.Lock()
	defer s.Unlock()

	s.hash.Reset()
	s.hash.Write([]byte(strToSign))
	sha := hex.EncodeToString(s.hash.Sum(nil))
	return sha
}
