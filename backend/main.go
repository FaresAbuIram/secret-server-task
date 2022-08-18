package main

import (
	"log"
	"secret-server-task/backend/database"
	"secret-server-task/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.Connect()

	router := gin.Default()

	router.Use(cors.Default())
	routes.Setup(router)
	router.Run("localhost:9090")
}
