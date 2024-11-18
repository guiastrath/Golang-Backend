package common

// Recognition Data
const (
	RecognitionURL         string = "http://localhost:8000/api/v1/recognition/recognize"
	RecognitionToken       string = "b90d59f0-5fa0-4222-a8b5-4258d220caeb"
	RecognitionContentType string = "multipart/form-data"
)

// Validation Data
const (
	ValidationToken string = "b90d59f0-5fa0-4222-a8b5-4258d220caeb"
)

// Barrier Data
const (
	BarrierControlUrl     string = "http://168.168.52.19:5555/zeroconf/switch"
	BarrierPulseConfigUrl string = "http://168.168.52.19:5555/zeroconf/pulse"
	TestBarrierPresent    bool   = false

	DeviceId          string  = "b90d59f0-5fa0-4222-a8b5-4258d220caeb"
	MinimumSimilarity float32 = 0.85
)

// Barrier Controls
const (
	ON  string = "on"
	OFF string = "off"
)

// Content Types
const (
	JSON string = "application/json"
)
