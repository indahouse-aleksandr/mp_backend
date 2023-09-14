package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("START")
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/time", TimeNow)
	http.ListenAndServe(":80", nil)
	fmt.Println("END")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Healthz, %s!", r.URL.Path[1:])
}

func TimeNow(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	t := time.Now()
	fmt.Fprintf(w, "TIME: , %s!", t)
}
