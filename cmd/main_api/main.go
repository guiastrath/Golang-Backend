package main

import (
	"golang-backend/internal/api/v1/handlers"
	"golang-backend/internal/middleware/auth"
	"golang-backend/pkg/httprest"
)

const (
	PORT = ":8001"
)

func main() {
	configs := &httprest.WSConfig{
		ServerPort: PORT,
	}

	server := httprest.NewWebService(configs)

	server.UseAuth(auth.NewAuthMiddleware).
		AddHandlers(
			handlers.NewRecognitionHandler(),
			handlers.NewConfigsHandler(),
		)

	server.ListenAndServe()
}
