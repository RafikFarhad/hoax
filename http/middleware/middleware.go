package middleware

import (
	"github.com/RafikFarhad/hoax/config"
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
	utils.Swagger(http)
	// Default middlewares
	http.Use(recover.New())
	http.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	http.Use(helmet.New())
	http.Use(cors.New())
	// Rate Limiter
	// http.Use(limiter.New())

	// Logger & pprof
	if config.Debug {
		http.Use(logger.New())
		http.Use(pprof.New())
	} else {
		http.Use(etag.New())
	}
}
