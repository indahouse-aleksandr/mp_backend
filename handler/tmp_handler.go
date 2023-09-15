package handler

import (
	"fmt"
	"net/http"
	"tc_backend/service"
)

func TmpHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	srv := service.NewTmpService()
	answer := srv.GetAnswer(r.Context())

	res := ""
	for _, v := range answer {
		fmt.Println(v)
		res += fmt.Sprintf(v.Title.String)
	}

	fmt.Fprintf(w, res)
}
