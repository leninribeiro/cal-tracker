package data

import (
	"context"
	"fmt"

	"github.com/leninribeiro/cal-tracker/config"
	"github.com/leninribeiro/cal-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection = Client.Database(config.C.DBName).Collection("users")

func InsertOneUser(user models.User) (*mongo.InsertOneResult, error) {
	result, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return result, fmt.Errorf("could not insert user '%s': %s", user.Name, err)
	}
	return result, err
}

func FindUser(userFilter bson.M) ([]models.User, error) {
	var users []models.User
	cursor, err := usersCollection.Find(context.TODO(), userFilter)

	if err != nil {
		return users, fmt.Errorf("error when querying mongodb: %s", err)
	}

	if err := cursor.All(context.TODO(), &users); err != nil {
		return users, fmt.Errorf("error when unmarshaling mongodb cursor: %s", err)
	}

	return users, err
}
