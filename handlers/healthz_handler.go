package handlers

import (
	"fmt"
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthz, %s!", r.URL.Path[1:])
}
