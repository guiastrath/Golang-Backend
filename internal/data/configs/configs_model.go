package configs

type PulseConfigData struct {
	Pulse      string `json:"pulse"`
	PulseWidth int    `json:"pulseWidth"`
}

type BarrierConfig struct {
	DeviceID string          `json:"deviceid"`
	Data     PulseConfigData `json:"data"`
}
