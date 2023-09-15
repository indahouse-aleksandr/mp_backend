package handler

import (
	"fmt"
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	fmt.Fprintf(w, "Healthz, %s!", r.URL.Path[1:])
}
