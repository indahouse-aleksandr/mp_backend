package middleware

import "net/http"

func Public(handler http.HandlerFunc) http.HandlerFunc {
	return CORS(handler)
}

func Private(handler http.HandlerFunc) http.HandlerFunc {
	return JWT(CORS(handler))
}
