package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DbInstance Function will return a Mongo Client which will help us to connect to Db and interact with Collections
func DbInstance() *mongo.Client {
	MongoDb := "mongo://localhost:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil{
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to MongoDb.")

	return client
}

var Client *mongo.Client = DbInstance()

// OpenCollection will provide us access to Collectin collectionName.
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("restaurant").Collection(collectionName)
}