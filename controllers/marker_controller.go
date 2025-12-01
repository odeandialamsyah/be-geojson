package controllers

import (
	"be-geojson/config"
	"be-geojson/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var markerCollection *mongo.Collection

func InitMarkerCollection() {
    markerCollection = config.Mongo.Collection("markers")
}


func GetMarkers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := markerCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var markers []models.Marker
	if err := cursor.All(ctx, &markers); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(markers)
}

func CreateMarker(c *fiber.Ctx) error {
	var marker models.Marker

	if err := c.BodyParser(&marker); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	marker.Location.Type = "Point"

	_, err := markerCollection.InsertOne(ctx, marker)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(marker)
}
