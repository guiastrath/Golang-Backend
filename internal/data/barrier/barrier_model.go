package barrier

type BarrierControlData struct {
	Switch string `json:"switch"`
}

type BarrierControl struct {
	DeviceID string             `json:"deviceid"`
	Data     BarrierControlData `json:"data"`
}
