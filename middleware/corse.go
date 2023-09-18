package middleware

import (
	"fmt"
	"net/http"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("CORS added")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("CORS processing ...")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
