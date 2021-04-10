package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func JsonResponse(c *fiber.Ctx) error {
	return c.Type("json", "utf-8").Next()
}
