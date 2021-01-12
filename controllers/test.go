package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Tasks struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []*Tasks{
	{
		Id:        1,
		Title:     "Create API",
		Completed: true,
	},
	{
		Id:        2,
		Title:     "Update API",
		Completed: false,
	},
}

func GetTasks(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"tasks": tasks,
		},
	})
}

func CreateTask(c *fiber.Ctx) error {
	type Request struct {
		Title string `json:"title"`
	}

	var body Request

	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON.",
		})
	}

	task := &Tasks{
		Id:        len(tasks) + 1,
		Title:     body.Title,
		Completed: false,
	}

	tasks = append(tasks, task)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"task": task,
		},
	})
}

func GetTask(c *fiber.Ctx) error {
	paramId := c.Params("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse ID.",
		})
	}

	for _, task := range tasks {
		if task.Id == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"task": task,
				},
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Task not found",
	})

}
