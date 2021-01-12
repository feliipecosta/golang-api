package routes

import (
	"github.com/feliipecosta/golang-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func TasksRoute(route fiber.Router) {
	route.Get("", controllers.GetTasks)
	route.Post("", controllers.CreateTask)
	// route.Put("/:id", controllers.UpdateTask)
	// route.Delete("/:id", controllers.DeleteTask)
	route.Get("/:id", controllers.GetTask)

}
