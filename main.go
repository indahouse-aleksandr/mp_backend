package main

import (
	"github.com/gin-gonic/gin"
	"github.com/indahouse-aleksandr/tc_backend/routers"
)

func main() {
	r := gin.Default()
	r = routers.SetupRouter(r)
	r.Run(":80")
}
