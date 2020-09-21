package account

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type GetMmpResponse struct {
	base.RestBaseResponse
	Data MmpStateVo `json:"data"`
}

type MmpStateVo struct {
	Currency            string      `json:"currency"`
	MmpEnabled          bool        `json:"mmp_enabled"`
	MmpUserConfigurable bool        `json:"mmp_user_configurable"`
	MmpConfig           MmpConfigVo `json:"mmp_config"`
	MmpFrozenUntilMs    int64       `json:"mmp_frozen_until_ms"`
	MmpFrozen           bool        `json:"mmp_frozen"`
}

type MmpConfigVo struct {
	WindowMs       int64  `json:"window_ms"`
	FrozenPeriodMs int64  `json:"frozen_period_ms"`
	QtyLimit       string `json:"qty_limit"`
	DeltaLimit     string `json:"delta_limit"`
}
