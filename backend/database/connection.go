package database

import (
	"context"
	"log"
	"os"
	"time"

	"secret-server-task/backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	SecretsCollection *mongo.Collection
	Ctx               = context.TODO()
)

func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal(err)
	}

	Ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(Ctx)
	if err != nil {
		log.Fatal(err)
	}

	secretDatabase := client.Database(os.Getenv("MONGO_DATABASE"))
	SecretsCollection = secretDatabase.Collection(os.Getenv("MONGO_COLLECTION"))
}

func GetSecret(id string) (models.Secret, error) {
	var s models.Secret
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
		return s, err
	}

	err = SecretsCollection.
		FindOne(Ctx, bson.D{{Key: "_id", Value: objectId}}).
		Decode(&s)
	if err != nil {
		log.Fatal(err)
		return s, err
	}
	return s, nil
}

func UpdateSecrete(id primitive.ObjectID, views int) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "views", Value: views}}}}
	_, err := SecretsCollection.UpdateOne(
		Ctx,
		filter,
		update,
	)
	return err
}

func DeleteSecret(id primitive.ObjectID) error {
	_, err := SecretsCollection.DeleteOne(Ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	return nil
}
