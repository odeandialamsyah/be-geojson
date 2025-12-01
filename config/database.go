package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database

func ConnectMongo() {
	// Ambil URI dari .env
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		// fallback default (kompatibel docker-compose)
		uri = "mongodb://mongo:27017/gisdb"
	}

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("❌ Error creating Mongo client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("❌ Mongo Connect Error:", err)
	}

	// Ambil nama database dari URI
	Mongo = client.Database("gisdb")

	log.Println("✅ MongoDB Connected:", uri)
}
