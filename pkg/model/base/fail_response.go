package base

type FailResponse struct {
	RestBaseResponse
	Data interface{} `json:"data"`
}
