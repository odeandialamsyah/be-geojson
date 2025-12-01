package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database

func ConnectMongo() {
	godotenv.Load(".env")
	// Ambil URI dari ENV
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("❌ MONGO_URI is not set in environment variables")
	}

	// Setup client
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("❌ Failed to create Mongo client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}

	// Ambil nama DB (default: gisdb)
	Mongo = client.Database("gisdb")

	log.Println("✅ MongoDB Connected:", uri)
}
