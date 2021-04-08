package routes

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/RafikFarhad/hoax/http"
	"github.com/RafikFarhad/hoax/http/middleware"
)

func InitRoutes() {
	config := config.AppConfig
	httpApp := http.AppHttp
	// Register default middlewares
	middleware.DefaultMiddlewares(httpApp, config)

	// Register web routes
	InitWebRoutes(httpApp, config)

	// Register api_v1 routes
	InitApiRoutes(httpApp, config)
}
