package controllers

import "github.com/gofiber/fiber/v2"

type Test struct {
	Id        int    `json:id`
	Title     string `json:title`
	Completed bool   `json:completed`
}

var tests = []*Test{
	{
		Id:        1,
		Title:     "Testing API",
		Completed: false,
	},
	{
		Id:        2,
		Title:     "Second one",
		Completed: false,
	},
}

func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"tests": tests,
		},
	})
}
