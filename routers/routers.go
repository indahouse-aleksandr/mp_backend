package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/indahouse-aleksandr/mp_backend/controllers"
)

func SetupRouter(engine *gin.Engine) *gin.Engine {

	engine.GET("/ping", controllers.Ping)
	engine.GET("/user/:name", controllers.UserName)

	return engine
}
