package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xolisamatika/go-rest-api/internal/api/routes"
	"github.com/xolisamatika/go-rest-api/internal/pkg/config"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
}

func setupRoutes(app *fiber.App) {

	api := app.Group("/api")
	routes.DirectoryRoute(api.Group("/directories"))
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	err := app.Listen(":" + conf.Server.Port)

	if err != nil {
		panic(err)
	}
}
