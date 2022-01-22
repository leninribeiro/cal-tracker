package data

import (
	"context"
	"fmt"
	"log"

	"github.com/leninribeiro/cal-tracker/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func LoadMongo() {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.C.DBName, config.C.DBPass, config.C.DBHost, config.C.DBPort)
	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Printf("Could not access database: %s", err)
	}
	fmt.Printf("Successfuly connected to MongoDB: %s", mongoURI)
}
