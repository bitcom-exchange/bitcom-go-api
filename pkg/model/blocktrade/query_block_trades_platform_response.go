package blocktrade

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type QueryBlockTradeByPlatformResponse struct {
	base.RestBaseResponse
	Data []ParadigmBlockTradeQueryVo `json:"data"`
}
