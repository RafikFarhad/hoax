package http

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/gofiber/fiber/v2"
)

var AppHttp *fiber.App

func CreateHTTPServer() error {
	appConfig := config.AppConfig
	AppHttp = fiber.New(fiber.Config{
		Prefork: appConfig.Prefork,
	})
	return nil
}
