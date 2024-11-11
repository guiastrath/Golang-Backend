package httprest

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrJSONEncoder       = errors.New("error encoding data")
	ErrReadBody          = errors.New("error reading request body")
	ErrJSONUnmarshal     = errors.New("JSON structure error")
	ErrInvalidStruct     = errors.New("provided type is not a valid struct")
	ErrMissingQueryParam = errors.New("missing query parameter")
)

type ErrorResponse struct {
	StatusCode int
	Message    string
}

func JSON(w http.ResponseWriter, statusCode int, jsonData []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if jsonData != nil {
		if _, err := w.Write(jsonData); err != nil {
			writeError(w, http.StatusInternalServerError, ErrJSONEncoder)
			return
		}
	}
}

func Response(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			writeError(w, http.StatusInternalServerError, ErrJSONEncoder)
			return
		}
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	writeError(w, statusCode, err)
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	response := ErrorResponse{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
	Response(w, statusCode, response)
}
