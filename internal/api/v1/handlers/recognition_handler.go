package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"golang-backend/internal/data/barrier"
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
	barrierService     barrier.BarrierService
}

func NewRecognitionHandler() *RecognitionHandler {
	return &RecognitionHandler{
		recognitionService: recognition.NewRecognitionService(),
		barrierService:     barrier.NewBarrierService(),
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
	sessionInfo := r.Context().Value(auth.AUTH_DATA).(*auth.AuthData)
	if sessionInfo == nil {
		fmt.Println("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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

	response, err := h.recognitionService.Recognize(sessionInfo, fileData, params)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	validatedFace, err := h.barrierService.ControlBarrier(sessionInfo, response)

	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("barrier control failed: %v", err), http.StatusInternalServerError)
		return
	}

	if validatedFace == nil {
		httprest.Response(w, http.StatusOK, "access denied")
		return
	}

	httprest.Response(w, http.StatusOK, validatedFace)
}
