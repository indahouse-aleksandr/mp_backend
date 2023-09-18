package handlers

import (
	"fmt"
	"net/http"
)

func AuthPage(w http.ResponseWriter, r *http.Request) {

	res := `<a target="_blank" href="/auth-get-google">get google</a>`

	fmt.Fprintf(w, res)
}
