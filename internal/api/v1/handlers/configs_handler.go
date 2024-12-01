package handlers

import (
	"fmt"
	"golang-backend/internal/data/barrier"
	"golang-backend/internal/middleware/auth"
	"golang-backend/pkg/httprest"
	"net/http"
)

const (
	configsBaseUrl = "/configs"
)

type ConfigsHandler struct {
	barrierService barrier.BarrierService
}

func NewConfigsHandler() *ConfigsHandler {
	return &ConfigsHandler{
		barrierService: barrier.NewBarrierService(),
	}
}

func (h *ConfigsHandler) Handlers() []*httprest.Route {
	return httprest.PrivateRoutes("/api/v1",
		httprest.GET(configsBaseUrl+"/barrier").To(h.ConfigureBarrier),
	)
}

func (h *ConfigsHandler) ConfigureBarrier(w http.ResponseWriter, r *http.Request) {
	sessionInfo := r.Context().Value(auth.AUTH_DATA).(*auth.AuthData)
	if sessionInfo == nil {
		fmt.Println("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	h.barrierService.ConfigureBarrier(sessionInfo, "on", 2000)
}
