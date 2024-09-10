package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var (
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
)

func Connect() error {
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}

	database = client.Database("go-mongodb")
	collection = database.Collection("user")
	fmt.Println("Connected to MongoDB")

	return nil
}

func Disconnect() {
	if client != nil {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Disconnected from MongoDB")
}

func Collection() *mongo.Collection {
	return collection
}
