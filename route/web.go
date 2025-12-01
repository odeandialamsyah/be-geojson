package routes

import (
	"be-geojson/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/markers", controllers.GetMarkers)
	api.Post("/markers", controllers.CreateMarker)
	api.Put("/markers/:id", controllers.UpdateMarker)
	api.Delete("/markers/:id", controllers.DeleteMarker)

	api.Get("/areas", controllers.GetAreas)
	api.Post("/areas", controllers.CreateArea)
	api.Put("/areas/:id", controllers.UpdateArea)
	api.Delete("/areas/:id", controllers.DeleteArea)
}
