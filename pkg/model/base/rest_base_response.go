package base

type RestBaseResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type OldPaging struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type NewPaging struct {
	HasMore bool `json:"has_more"`
}
