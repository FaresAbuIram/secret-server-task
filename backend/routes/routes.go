package routes

import (
	"secret-server-task/backend/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.POST("/generate", controllers.GenerateToken)
	router.GET("/get/:token", controllers.GetToken)
}
