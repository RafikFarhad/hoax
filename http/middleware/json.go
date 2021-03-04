package middleware

import "github.com/gofiber/fiber/v2"

func JsonResponse(c *fiber.Ctx) error {
	c.Type("json", "utf-8")
	return c.Next()
}
