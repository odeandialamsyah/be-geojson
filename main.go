package main

import (
	"be-geojson/config"
	routes "be-geojson/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectMongo()

	routes.SetupRoutes(app)

	app.Listen(":8080")
}
