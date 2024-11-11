package auth

import (
	"context"
	"errors"
	"golang-backend/pkg/common"
	"golang-backend/pkg/httprest"
	"net/http"
)

type Token string

const (
	API_KEY Token = "apiKey"
)

func NewAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("x-api-key")

		if headerToken == "" {
			httprest.Error(w, http.StatusUnauthorized, errors.New("token is missing"))
			return
		}

		if headerToken != common.RecognitionToken {
			httprest.Error(w, http.StatusUnauthorized, errors.New("unauthorized, invalid token"))
			return
		}

		ctx := context.WithValue(r.Context(), API_KEY, headerToken)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
