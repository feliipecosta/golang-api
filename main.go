package main

import (
	"github.com/feliipecosta/golang-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Here is an endpoint.",
		})
	})

	api := app.Group("/api")

	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "API endpoint.",
		})
	})

	routes.TasksRoute(api.Group("/tasks"))
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// setting up routes
	setupRoutes(app)

	// Listen to port 3000
	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}

}
