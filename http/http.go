package http

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/gofiber/fiber/v2"
)

func CreateHTTPServer(config *config.HoaxConfig) (*fiber.App, error) {

	return fiber.New(fiber.Config{
		Prefork: config.Prefork,
	}), nil
}
