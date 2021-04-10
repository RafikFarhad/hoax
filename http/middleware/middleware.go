package middleware

import (
	"github.com/RafikFarhad/hoax/config"
	logger2 "github.com/RafikFarhad/hoax/logger"
	"github.com/RafikFarhad/hoax/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

func DefaultMiddlewares(http *fiber.App, config *config.HoaxConfig) {
	// swagger setup
	if config.SwaggerEnabled {
		utils.Swagger(http)
	}

	// default middlewares
	http.Use(recover.New())
	http.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	http.Use(helmet.New())
	http.Use(cors.New())

	// rate limiter
	// http.Use(limiter.New())

	// logger & pprof
	if config.Debug {
		http.Use(logger.New(logger.Config{
			Output: logger2.GlobalLogger.Writer(),
		}))
		http.Use(pprof.New())
	} else {
		http.Use(etag.New())
	}
}
