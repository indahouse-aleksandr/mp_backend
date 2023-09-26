package main

import (
	"github.com/gin-gonic/gin"
	"github.com/indahouse-aleksandr/mp_backend/application"
	"github.com/indahouse-aleksandr/mp_backend/validators"
)

func main() {
	r := gin.Default()
	r = application.SetupRouter(r)
	r = validators.RegisterCustomValidators(r)

	r.Run(":80")
}
