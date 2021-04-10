package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Welcome(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("Hello World!!!")
}

func About(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("This is about !!!")

}
