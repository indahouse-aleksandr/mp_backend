package handler

import (
	"fmt"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
