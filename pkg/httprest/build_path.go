package httprest

import (
	"fmt"
	"net/http"
)

func GET(path string) string {
	return fmt.Sprintf("%s %s", http.MethodGet, path)
}

func POST(path string) string {
	return fmt.Sprintf("%s %s", http.MethodPost, path)
}

func DELETE(path string) string {
	return fmt.Sprintf("%s %s", http.MethodDelete, path)
}

func PATCH(path string) string {
	return fmt.Sprintf("%s %s", http.MethodPatch, path)
}

func PUT(path string) string {
	return fmt.Sprintf("%s %s", http.MethodPut, path)
}
