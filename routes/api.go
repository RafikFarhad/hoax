package routes

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/RafikFarhad/hoax/http/controllers/api_v1"
	"github.com/RafikFarhad/hoax/http/controllers/api_v1/auth"
	"github.com/RafikFarhad/hoax/http/middleware"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func InitApiRoutes(http *fiber.App, config *config.HoaxConfig) {
	apiHttp := http.Group("api/v1", middleware.JsonResponse)

	// Basic
	apiHttp.Get("ping", api_v1_basic.Ping)

	// Guest
	apiHttp.Post("login", api_v1_auth.Login)

	// Auth
	if config.JwtSecret == "" {
		config.JwtSecret = "random"
	}
	authApi := apiHttp.Group("", jwtware.New(jwtware.Config{
		SigningKey:     []byte(config.JwtSecret),
		SuccessHandler: middleware.AuthSuccessHandler,
		ErrorHandler:   middleware.AuthErrorHandler,
		TokenLookup:    "header:Authorization",
		AuthScheme:     "Bearer",
		ContextKey:     "token_data",
	}))

	authApi.Get("me", api_v1_auth.Me)
}
