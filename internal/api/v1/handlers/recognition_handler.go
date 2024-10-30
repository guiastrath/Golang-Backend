package handlers

import (
	"golang-backend/internal/data/recognition"
	"golang-backend/internal/middleware/auth"
	"golang-backend/pkg/httprest"
	"net/http"
)

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

func (h *RecognitionHandler) BuildHandlers(mux *http.ServeMux) {
	mux.HandleFunc(httprest.GET(recognitionBaseUrl+"/hello-world"), h.HelloWorld)
	mux.HandleFunc(httprest.GET(recognitionBaseUrl+"/recognize"), h.Recognize)
}

func (h *RecognitionHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *RecognitionHandler) Recognize(w http.ResponseWriter, r *http.Request) {
	apiKey, ok := r.Context().Value(auth.ApiKey).(string)

	if !ok || apiKey != auth.Token {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// ctx := r.Context()

	// recognition, err := h.recognitionService.Recognize(ctx, apiKey)

	// if err != nil {
	// 	http.Error(w, "Error", http.StatusInternalServerError)
	// 	return
	// }

	httprest.Response(w, http.StatusOK, "recognition")
}
