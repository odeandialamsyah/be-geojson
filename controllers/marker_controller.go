package controllers

import (
	"be-geojson/config"
	"be-geojson/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	res, err := markerCollection.InsertOne(ctx, marker)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// set ID if returned
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		marker.ID = oid
	}

	return c.JSON(marker)
}

func UpdateMarker(c *fiber.Ctx) error {
	idParam := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString("invalid id")
	}

	var marker models.Marker
	if err := c.BodyParser(&marker); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	marker.Location.Type = "Point"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"name": marker.Name, "location": marker.Location}}
	_, err = markerCollection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	marker.ID = oid
	return c.JSON(marker)
}

func DeleteMarker(c *fiber.Ctx) error {
	idParam := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString("invalid id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = markerCollection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
