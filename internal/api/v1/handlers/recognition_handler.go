package handlers

import (
	"fmt"
	"golang-backend/internal/data/recognition"
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

func (h *RecognitionHandler) Handlers() []*httprest.Route {
	return httprest.PrivateRoutes("/v1",
		httprest.GET(recognitionBaseUrl+"/helloworld").To(h.HelloWorld),
		httprest.POST(recognitionBaseUrl+"/recognize").To(h.Recognize),
	)
}

func (h *RecognitionHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *RecognitionHandler) Recognize(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var p []byte
	a, err := h.recognitionService.Recognize(ctx, r.Body)
	r.Body.Read(p)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	fmt.Println(p)
	// apiKey, ok := r.Context().Value(auth.API_KEY).(string)

	// if !ok || apiKey != auth.RecognitionToken {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	// ctx := r.Context()

	// recognition, err := h.recognitionService.Recognize(ctx, apiKey)

	httprest.Response(w, http.StatusOK, a)
}
