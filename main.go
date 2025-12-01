package main

import (
	"be-geojson/config"
	routes "be-geojson/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization",
		AllowMethods:  "GET,POST,PUT,DELETE,OPTIONS",
	}))

	config.ConnectMongo()

	routes.SetupRoutes(app)

	app.Listen(":8080")
}
