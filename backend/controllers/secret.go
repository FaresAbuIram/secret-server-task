package controllers

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"secret-server-task/backend/database"
	"secret-server-task/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateToken(context *gin.Context) {
	var input models.Secret
	if err := context.BindJSON(&input); err != nil {
		fmt.Println("err: ", err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on create secret.", "data": err})
		return
	}
	result, err := database.SecretsCollection.InsertOne(database.Ctx, bson.D{
		{Key: "expire", Value: input.Expire},
		{Key: "views", Value: input.Views},
	})
	var mySigningKey = []byte(os.Getenv("SECRET_TOKEN"))
	expirationTime := time.Now().Add(time.Hour * time.Duration(input.Expire))
	fmt.Println(expirationTime)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = result.InsertedID
	claims["data"] = input.Data
	claims["exp"] = expirationTime.Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println(err)

	}

	context.JSON(http.StatusOK, gin.H{"data": tokenString})

}

func GetToken(context *gin.Context) {
	tokenString := context.Param("token")

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_TOKEN")), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	info := token.Claims.(*jwt.MapClaims)
	fmt.Println("info: ", (*info)["id"])
	context.JSON(http.StatusOK, gin.H{"data": token.Claims})
}
