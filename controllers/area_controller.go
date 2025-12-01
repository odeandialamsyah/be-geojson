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

var areaCollection *mongo.Collection

func InitAreaCollection() {
    areaCollection = config.Mongo.Collection("areas")
}


func GetAreas(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := areaCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var areas []models.Area
	if err := cursor.All(ctx, &areas); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(areas)
}

func CreateArea(c *fiber.Ctx) error {
	var area models.Area

	if err := c.BodyParser(&area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	area.Coordinates.Type = "Polygon"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := areaCollection.InsertOne(ctx, area)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		area.ID = oid
	}

	return c.JSON(area)
}

func UpdateArea(c *fiber.Ctx) error {
	idParam := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString("invalid id")
	}

	var area models.Area
	if err := c.BodyParser(&area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	area.Coordinates.Type = "Polygon"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"name": area.Name, "coordinates": area.Coordinates}}
	_, err = areaCollection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	area.ID = oid
	return c.JSON(area)
}

func DeleteArea(c *fiber.Ctx) error {
	idParam := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString("invalid id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = areaCollection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
