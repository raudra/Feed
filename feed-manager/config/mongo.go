package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitMongo() {
	log.Println("Setting up mongo client")
	mongiClient, err := mongo.NewClient(options.
		Client().
		ApplyURI("mongodb://host.docker.internal:27017"))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongiClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = mongiClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
}

func MongoClient() *mongo.Client {
	return mongoClient
}
