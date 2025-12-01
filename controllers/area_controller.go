package controllers

import (
	"be-geojson/config"
	"be-geojson/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var areaCollection = config.Mongo.Collection("areas")

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

	_, err := areaCollection.InsertOne(ctx, area)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(area)
}
