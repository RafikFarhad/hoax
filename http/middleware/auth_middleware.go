package middleware

import (
	"github.com/RafikFarhad/hoax/http/response"
	"github.com/RafikFarhad/hoax/types"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthErrorHandler(c *fiber.Ctx, err error) error {
	return response.WithError(c, "AUTH_ERROR :- "+err.Error(), 1, 401)
}

func AuthSuccessHandler(c *fiber.Ctx) error {
	authToken := c.Locals("token_data").(*jwt.Token)
	claims := authToken.Claims.(jwt.MapClaims)
	c.Locals("auth_data", types.AuthData{
		UserId: uint(claims["user_id"].(float64)),
	})
	return c.Next()
}
