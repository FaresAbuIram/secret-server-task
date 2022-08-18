package controllers

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"secret-server-task/backend/database"
	"secret-server-task/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Generate new secret
// @Summary      Create a new secret
// @Description  This route generates a new secret have the user's data
// @Tags         token
// @Accept       json
// @Produce      json
// @Param        body  body      models.Data  true  "Generate a secret"
// @Success      200  {object}  models.ResultToken
// @Failure      400  {object}	  object
// @Router       /generate [post]
func GenerateToken(context *gin.Context) {
	var input models.Data
	if err := context.BindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error in creating a secret."})
		return
	}

	var mySigningKey = []byte(os.Getenv("SECRET_TOKEN"))
	expirationTime := time.Now().Add(time.Minute * time.Duration(input.Expire))

	// add object to database
	result, err := database.SecretsCollection.InsertOne(database.Ctx, bson.D{
		{Key: "expire_date", Value: expirationTime},
		{Key: "views", Value: input.Views},
	})
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error in creating a secret."})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = result.InsertedID
	claims["data"] = input.Data
	claims["exp"] = expirationTime.Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error in creating a secret."})
	}

	context.JSON(http.StatusOK, gin.H{"data": tokenString})

}

// Get the secret Information
// @Summary      analyze the secret
// @Description  This routes generate new secret have the user's data
// @Tags         token
// @Accept       json
// @Produce      json
// @Param        token  path      string  true  "get the secret info"
// @Success      200  {object}    models.ResponseData
// @Failure      400  {object}	  object
// @Failure      500  {object}	  object
// @Router       /get/{token} [post]
func GetToken(context *gin.Context) {
	tokenString := context.Param("token")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_TOKEN")), nil
	})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	info := token.Claims.(*jwt.MapClaims)
	id := (*info)["id"]
	secretData := (*info)["data"]
	objectId, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	// get the token object from database
	var object models.Secret
	err = database.SecretsCollection.
		FindOne(database.Ctx, bson.D{{Key: "_id", Value: objectId}}).
		Decode(&object)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	if object.Views <= 0 {
		// delete expired object
		_, err := database.SecretsCollection.DeleteOne(database.Ctx, bson.D{{Key: "_id", Value: objectId}})
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "something get wrong"})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": "No views available"})
		return
	}
	// update the number of views
	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "views", Value: object.Views - 1}}}}
	_, err = database.SecretsCollection.UpdateOne(database.Ctx, filter, update)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "something get wrong"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": secretData, "object": object, "message": "success"})
}
