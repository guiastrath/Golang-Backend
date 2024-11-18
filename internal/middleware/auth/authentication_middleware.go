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
	AUTH_DATA Token = "authData"
)

func NewAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("x-api-key")

		if headerToken == "" {
			httprest.Error(w, http.StatusBadRequest, errors.New("token is missing"))
			return
		}

		if headerToken != common.RecognitionToken {
			httprest.Error(w, http.StatusUnauthorized, errors.New("unauthorized, invalid token"))
			return
		}

		deviceId := r.Header.Get("deviceId")

		if deviceId == "" {
			httprest.Error(w, http.StatusBadRequest, errors.New("deviceId is missing"))
			return
		}

		if deviceId != common.DeviceId {
			httprest.Error(w, http.StatusBadRequest, errors.New("invalid deviceId"))
			return
		}

		ctx := context.WithValue(r.Context(), AUTH_DATA, &AuthData{
			ApiKey:   headerToken,
			DeviceID: deviceId,
		})
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
