package application

import (
	"github.com/gin-gonic/gin"
	"github.com/indahouse-aleksandr/mp_backend/controllers"
)

func SetupRouter(engine *gin.Engine) *gin.Engine {

	engine.GET("/user/:id", controllers.GetUserByID)
	engine.POST("/user", controllers.CreateUser)

	return engine
}
