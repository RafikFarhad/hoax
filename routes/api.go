package routes

import (
	"github.com/RafikFarhad/hoax/http/controllers/api_v1"
	"github.com/RafikFarhad/hoax/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitApiRoutes(http *fiber.App) {
	apiHttp := http.Group("api/v1", middleware.JsonResponse)
	apiHttp.Get("ping", api_v1_auth.Ping)
	apiHttp.Post("login", api_v1_auth.Login)
}
