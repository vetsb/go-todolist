package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	options := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, options)

	if err == nil {
		fmt.Println("Mongodb is connected")
	} else {
		fmt.Println("Mongodb is not connected")
		return
	}

	mongoClient = client

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func GetDefaultDatabase() *mongo.Database {
	return mongoClient.Database("todolist")
}
