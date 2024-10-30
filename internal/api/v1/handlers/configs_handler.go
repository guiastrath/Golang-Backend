package handlers

import (
	"golang-backend/pkg/httprest"
	"net/http"
)

const (
	configsBaseUrl = "/configs"
)

type ConfigsHandler struct {
	// configsService configs.ConfigsService
}

func NewConfigsHandler() *ConfigsHandler {
	return &ConfigsHandler{
		// configsService: configs.NewConfigsService(),
	}
}

func (h *ConfigsHandler) BuildHandlers(mux *http.ServeMux) {
	mux.HandleFunc(httprest.GET(configsBaseUrl+"/"), h.HelloWorld)
}

func (h *ConfigsHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
