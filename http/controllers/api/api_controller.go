package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Ping(ctx *fiber.Ctx) error {
	ctx.Send([]byte("{ \"Message\": \"PONG\"}"))
	return nil
}
