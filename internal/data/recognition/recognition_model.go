package recognition

type RecognitionRequest struct {
	File      string `json:"file"`
	BarrierID string `json:"barrierId"`
}

type Box struct {
	Probability float32 `json:"probability"`
	XMax        int     `json:"x_max"`
	YMax        int     `json:"y_max"`
	XMin        int     `json:"x_min"`
	YMin        int     `json:"y_min"`
}

type Subject struct {
	Subject    string  `json:"subject"`
	Similarity float32 `json:"similarity"`
}

type Result struct {
	Box      Box       `json:"box"`
	Subjects []Subject `json:"subjects"`
}

type RecognitionResponse struct {
	Result []Result `json:"result"`
}
