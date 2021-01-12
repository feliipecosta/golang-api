package routes

import (
	"github.com/feliipecosta/golang-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func TestRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
