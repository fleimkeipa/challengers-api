package pkg

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect() (*mongo.Database, error) {
	var uri = "mongodb://localhost:27017"
	if os.Getenv("STAGE") == "prod" {
		uri = "mongodb://mongodb:27017"
	}

	// Set client options
	var clientOptions = options.Client().ApplyURI(uri)
	clientOptions.SetDirect(true)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	// Set the database and collection variables
	return client.Database(os.Getenv("DATABASE_NAME")), nil
}
