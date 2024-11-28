package handlers

import (
	"golang-backend/internal/data/barrier"
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
		httprest.GET(configsBaseUrl+"/barrier/{deviceId}").To(h.ConfigureBarrier),
	)
}

func (h *ConfigsHandler) ConfigureBarrier(w http.ResponseWriter, r *http.Request) {
	// h.barrierService.ConfigureBarrier(deviceId, pulse, pulseWidth)
}
