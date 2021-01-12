package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "This is an endpoint",
		})
	})

	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}

}
