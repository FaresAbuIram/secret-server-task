package controllers

import (
	"fmt"
	"net/http"
	"time"

	"secret-server-task/backend/database"
	"secret-server-task/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
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
	var mySigningKey = []byte("secretkey")
	expirationTime := time.Now().Add(time.Microsecond * time.Duration(input.Expire))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = result
	claims["data"] = input.Data
	claims["exp"] = expirationTime

	
	tokenString, err := token.SignedString(mySigningKey)



	if err != nil {
		fmt.Println(err)
		
	}

	context.JSON(http.StatusOK, gin.H{"status": "success", "message": "User logged in.", "data": tokenString})

}

func GetToken(context *gin.Context) {

}
