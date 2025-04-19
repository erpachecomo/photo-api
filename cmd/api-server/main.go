package main

import (
	"context"
	"log"
	"os"

	"github.com/erpachecomo/photo-api/internal/api"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// Database connection.
	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalf("Error on db connection: %v", err)
	}

	defer func() {
		if err := client.Disconnect((context.TODO())); err != nil {
			panic(err)
		}
	}()
	db := client.Database("photo-api")
	api.SetupRoutes(db)

}
