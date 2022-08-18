package routes

import (
	"secret-server-task/backend/controllers"
	"secret-server-task/backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" 
)

func Setup(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9090"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.POST("/generate", controllers.GenerateToken)
	router.POST("/get/:token", controllers.GetToken)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
