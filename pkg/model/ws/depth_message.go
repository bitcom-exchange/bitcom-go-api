package ws

import "github.com/bitcom-exchange/bitcom-go-api/pkg/model/base"

type UpdateCtrlMessage struct {
	Type string
	Data interface{}
}

type DepthSnapshotResponse struct {
	base.WsBaseResponse
	Data DepthSnapshotMessage `json:"data"`
}

type DepthUpdateResponse struct {
	base.WsBaseResponse
	Data DepthUpdateMessage `json:"data"`
}

type Depth1Response struct {
	base.WsBaseResponse
	Data Depth1Message `json:"data"`
}

type DepthSnapshotMessage struct {
	Type         string      `json:"type,omitempty"`
	InstrumentId string      `json:"instrument_id"`
	Sequence     int64       `json:"sequence"`
	Bids         [][2]string `json:"bids"` // [[<price>, <qty>]]
	Asks         [][2]string `json:"asks"`
}

type Depth1Message struct {
	InstrumentId string      `json:"instrument_id"`
	Bids         [][2]string `json:"bids"` // [[<price>, <qty>]]
	Asks         [][2]string `json:"asks"`
}

type DepthUpdateMessage struct {
	Type         string      `json:"type"`
	InstrumentId string      `json:"instrument_id"`
	Sequence     int64       `json:"sequence"`
	PrevSequence int64       `json:"prev_sequence"`
	Changes      [][3]string `json:"changes"` // [[<side>, <price>, <qty>]]
}
