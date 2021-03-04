package middleware

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func DefaultMiddlewares(http *fiber.App, config *config.HoaxConfig) {

	// Default middlewares
	http.Use(recover.New())
	http.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	//// Logger & pprof
	if config.Debug {
		http.Use(logger.New())
		http.Use(pprof.New())
	} else {
		http.Use(etag.New())
	}
}
