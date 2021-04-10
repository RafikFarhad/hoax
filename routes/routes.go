package routes

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/RafikFarhad/hoax/http"
	"github.com/RafikFarhad/hoax/http/middleware"
)

func InitRoutes() {
	appConfig := config.AppConfig
	httpApp := http.AppHttp
	// Register default middlewares
	middleware.DefaultMiddlewares(httpApp, appConfig)

	// Register web routes
	InitWebRoutes(httpApp, appConfig)

	// Register api_v1 routes
	InitApiRoutes(httpApp, appConfig)
}
