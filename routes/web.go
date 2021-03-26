package routes

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/RafikFarhad/hoax/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitWebRoutes(http *fiber.App, _ *config.HoaxConfig) {

	http.Get("/", controllers.Welcome)

	http.Get("/about", controllers.About)
}
