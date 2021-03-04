package routes

import (
	controllers "github.com/RafikFarhad/hoax/http/controllers/api"
	"github.com/RafikFarhad/hoax/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitApiRoutes(http *fiber.App) {
	apiHttp := http.Group("api/v1", middleware.JsonResponse)
	apiHttp.Get("ping", controllers.Ping)
}
