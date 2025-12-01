package routes

import (
	"be-geojson/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/markers", controllers.GetMarkers)
	api.Post("/markers", controllers.CreateMarker)

	api.Get("/areas", controllers.GetAreas)
	api.Post("/areas", controllers.CreateArea)
}
