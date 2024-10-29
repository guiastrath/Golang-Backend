package handlers

import "golang-backend/internal/data/recognition"

const (
	recognitionBaseUrl = "/recognition"
)

type RecognitionHandler struct {
	recognitionService recognition.RecognitionService
}

func NewRecognitionHandler() *RecognitionHandler {
	return &RecognitionHandler{
		recognitionService: recognition.NewRecognitionService(),
	}
}

func (h *RecognitionHandler) Handlers() {

}
