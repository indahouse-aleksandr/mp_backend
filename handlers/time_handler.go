package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func TimeNow(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "TIME: , %s!", t)
}
