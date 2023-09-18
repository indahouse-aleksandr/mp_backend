package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func Uuid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid.New().String()
		ctx := context.WithValue(r.Context(), "RequestId", requestId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
