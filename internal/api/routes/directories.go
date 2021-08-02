package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xolisamatika/go-rest-api/internal/api/controllers"
)

func DirectoryRoute(route fiber.Router) {
	route.Get("", controllers.GetDirectoryList)
}
