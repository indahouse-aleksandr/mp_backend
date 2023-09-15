package main

import (
	"fmt"
	"net/http"
	"tc_backend/handler"
)

func main() {
	fmt.Println("START")

	//	fmt.Println(os.Getenv("DATABASE_URL"))
	//	os.Exit(2)

	http.HandleFunc("/", handler.HelloServer)
	http.HandleFunc("/healthz", handler.Healthz)
	http.HandleFunc("/time", handler.TimeNow)
	http.HandleFunc("/tmp", handler.TmpHandler)

	http.ListenAndServe(":80", nil)
	fmt.Println("END")
}
