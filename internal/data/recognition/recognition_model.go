package recognition

type RecognitionRequest struct {
	File string `json:"file"`
}

type Box struct {
	Probability float64 `json:"probability"`
	XMax        int     `json:"x_max"`
	YMax        int     `json:"y_max"`
	XMin        int     `json:"x_min"`
	YMin        int     `json:"y_min"`
}

type Subject struct {
	Subject    string  `json:"subject"`
	Similarity float64 `json:"similarity"`
}

type RecognitionResponse struct {
	Result []struct {
		Box      Box       `json:"box"`
		Subjects []Subject `json:"subjects"`
	} `json:"result"`
}
