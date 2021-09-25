package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var App *fiber.App

func Init() {
	App = fiber.New()

	App.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	setupRoutes()

	App.Listen(":8000")
}

func setupRoutes() {
	setupLinksRoutes()
}
