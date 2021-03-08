package routes

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/RafikFarhad/hoax/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(http *fiber.App, config *config.HoaxConfig) {
	// Register default middlewares
	middleware.DefaultMiddlewares(http, config)

	// Register web routes
	InitWebRoutes(http)

	// Register api_v1 routes
	InitApiRoutes(http)
}
