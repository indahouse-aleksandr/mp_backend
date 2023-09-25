package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/indahouse-aleksandr/tc_backend/controllers"
)

func SetupRouter(engine *gin.Engine) *gin.Engine {

	engine.GET("/ping", controllers.Ping)
	engine.GET("/user/:name", controllers.UserName)

	return engine
}
