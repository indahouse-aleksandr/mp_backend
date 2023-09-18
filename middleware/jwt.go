package middleware

import (
	"context"
	"fmt"
	"net/http"
)

func JWT(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("JWT added")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("JWT processing ...")
		UserId := 333
		ctx := context.WithValue(r.Context(), "UserId", UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
