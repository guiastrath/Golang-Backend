package main

import (
	"fmt"
	"golang-backend/internal/api/v1/handlers"
	"net/http"
)

const (
	PORT = ":8001"
)

func main() {

	mux := http.NewServeMux()

	// Add all Handlers
	handlers.NewRecognitionHandler().BuildHandlers(mux)
	handlers.NewConfigsHandler().BuildHandlers(mux)

	err := http.ListenAndServe(PORT, mux)
	fmt.Println(err)
	// fmt.Println("Hello, World!")
}
