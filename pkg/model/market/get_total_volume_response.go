package market

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetTotalVolumeResponse struct {
	base.RestBaseResponse
	Data TotalVolumeVo `json:"data"`
}

type TotalVolumeVo struct {
	TotalVolume24Hours string          `json:"total_volume_24_hours"`
	Details            []VolumeDetails `json:"details"`
}

type VolumeDetails struct {
	Currency      string `json:"currency"`
	Category      string `json:"category"`
	VolumeBaseCcy string `json:"volume_base_ccy"`
	VolumeUsd     string `json:"volume_usd"`
	IndexPrice    string `json:"index_price"`
}
