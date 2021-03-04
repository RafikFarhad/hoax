package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Welcome(ctx *fiber.Ctx) error {
	ctx.SendString("Hello World!!!")
	return ctx.SendStatus(200)
}

func About(ctx *fiber.Ctx) error {
	//App := app.App
	ctx.SendString("This is about !!!")
	return ctx.SendStatus(200)
}
