package middleware

import (
	"github.com/RafikFarhad/hoax/http/response"
	"github.com/RafikFarhad/hoax/types"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthErrorHandler(c *fiber.Ctx, err error) error {
	return response.ErrorMessageWithStatus(c, "AUTH_ERROR :- "+err.Error(), 401)
}

func AuthSuccessHandler(c *fiber.Ctx) error {
	authToken := c.Locals("token_data").(*jwt.Token)
	claims := authToken.Claims.(jwt.MapClaims)
	c.Locals("authData", types.AuthData{
		UserId: uint(claims["userId"].(float64)),
	})
	return c.Next()
}
