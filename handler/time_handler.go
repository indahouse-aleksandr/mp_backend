package handler

import (
	"fmt"
	"net/http"
	"time"
)

func TimeNow(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	t := time.Now()
	fmt.Fprintf(w, "TIME: , %s!", t)
}
