package routes

import (
	"github.com/RafikFarhad/hoax/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitWebRoutes(http *fiber.App) {

	http.Get("/", controllers.Welcome)

	http.Get("/about", controllers.About)
}
