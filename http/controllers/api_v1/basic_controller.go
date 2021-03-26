package api_v1_basic

import (
	"github.com/RafikFarhad/hoax/http/response"
	"github.com/gofiber/fiber/v2"
)

func Ping(ctx *fiber.Ctx) error {
	return response.SuccessMessage(ctx, "PONG")
}
