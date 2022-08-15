package main

import (
	"secret-server-task/backend/database"
	"secret-server-task/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	router := gin.Default()
	router.Use(cors.Default())
	routes.Setup(router)
	router.Run("localhost:9090")
}
