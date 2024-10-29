package main

import (
	"fmt"
	"net/http"
)

const (
	PORT = ":8001"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	err := http.ListenAndServe(PORT, mux)
	fmt.Println(err)
	// fmt.Println("Hello, World!")
}
