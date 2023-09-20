package config

import (
	"context"
	"fmt"
	"time"
	info "transaction1to10/infos"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDatabase() (*mongo.Client, error) {
	// Increase the timeout duration if needed
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second) // Example: 5 seconds
	defer cancel()

	// Create a MongoDB client with options
	clientOptions := options.Client().ApplyURI(info.Connectionstring)

	// Connect to the MongoDB server
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Check if the connection is established
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		// Close the client if the connection fails
		_ = mongoClient.Disconnect(ctx)
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	return mongoClient, nil
}
func Getcollection(dbName string, collectionName string) *mongo.Collection {
	client, err := ConnectDatabase()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Collection := client.Database(dbName).Collection(collectionName)
	return Collection
}
