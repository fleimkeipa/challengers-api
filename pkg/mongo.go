package pkg

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect() (*mongo.Database, error) {
	var uri = "mongodb://localhost:27017"
	if stage() {
		uri = "mongodb://mongodb:27017"
	}

	// Set client options
	var clientOptions = options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	// Set the database and collection variables
	return client.Database(os.Getenv("DATABASE_NAME")), nil
}

func stage() bool {
	_, isVirtual := os.LookupEnv("HOSTNAME")
	fmt.Printf("\nProgram is running inside a Docker container. %v\n", isVirtual)
	return isVirtual
}
