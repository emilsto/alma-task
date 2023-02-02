package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client instance
var Client *mongo.Client = ConnectDB()

// Connect to the database
func ConnectDB() *mongo.Client {

	conn, err := mongo.NewClient(options.Client().ApplyURI(GetConfig().MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// Add a timeout to the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = conn.Connect(ctx)
	if err != nil {
		log.Fatal("Timeout connecting to MongoDB", err)
	}

	// Check the connection
	err = conn.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB", err)
	}
	fmt.Println("Connected to MongoDB")

	// Return the connection
	return conn
}

// Getting the collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("beverages").Collection(collectionName)
	return collection
}
