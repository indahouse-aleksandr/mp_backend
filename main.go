package main

import (
	"fmt"
	"net/http"
	"tc_backend/handlers"
	"tc_backend/middleware"
)

func main() {
	fmt.Println("START")

	http.HandleFunc("/", middleware.Public(handlers.HelloServer))
	http.HandleFunc("/tmp", middleware.Private(handlers.TmpHandler))
	http.HandleFunc("/healthz", handlers.Healthz)
	http.HandleFunc("/time", handlers.TimeNow)

	http.HandleFunc("/auth-page", middleware.Public(handlers.AuthPage))
	http.HandleFunc("/auth-get-google", middleware.Public(handlers.AuthGetGoogle))
	http.HandleFunc("/auth", middleware.Public(handlers.AuthSetGoogle))

	http.ListenAndServe(":80", nil)
	fmt.Println("END")
}
