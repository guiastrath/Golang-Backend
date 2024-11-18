package barrier

type PulseConfigData struct {
	Pulse      string `json:"pulse"`
	PulseWidth int    `json:"pulseWidth"`
}

type BarrierConfig struct {
	DeviceID string          `json:"deviceid"`
	Data     PulseConfigData `json:"data"`
}

type BarrierControlData struct {
	Switch string `json:"switch"`
}

type BarrierControl struct {
	DeviceID string             `json:"deviceid"`
	Data     BarrierControlData `json:"data"`
}
