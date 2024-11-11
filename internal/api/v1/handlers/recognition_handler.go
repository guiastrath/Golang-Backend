package handlers

import (
	"encoding/base64"
	"encoding/json"
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
	return httprest.PrivateRoutes("/api/v1",
		httprest.GET(recognitionBaseUrl+"/helloworld").To(h.HelloWorld),
		httprest.POST(recognitionBaseUrl+"/recognize").To(h.Recognize),
	)
}

func (h *RecognitionHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *RecognitionHandler) Recognize(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(w, "No body.", http.StatusBadRequest)
	}

	params := r.URL.Query()

	request := &recognition.RecognitionRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	fileData, err := base64.StdEncoding.DecodeString(request.File)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding base64 data: %v", err), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	res, err := h.recognitionService.Recognize(ctx, fileData, params)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	httprest.Response(w, http.StatusOK, res)
}
