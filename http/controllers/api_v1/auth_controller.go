package api_v1_auth

import (
	"github.com/RafikFarhad/hoax/database/model"
	"github.com/RafikFarhad/hoax/http/response"
	"github.com/RafikFarhad/hoax/http/validator/api"
	"github.com/gofiber/fiber/v2"
)

func Ping(ctx *fiber.Ctx) error {
	return response.SuccessMessage(ctx, "PONG")
}

func Login(ctx *fiber.Ctx) error {
	loginRequest := &api_validator.LoginRequest{}
	// Parse
	if err := ctx.BodyParser(loginRequest); err != nil {
		return response.ErrorMessageWithStatus(ctx, err.Error(), 422)
	}
	// Validate
	if err := api_validator.ValidateLoginRequest(loginRequest); err != nil {
		return response.ValidationError(ctx, err)
	}
	// Verify
	if _, err := model.GetUserByUsername(loginRequest.Username); err != nil {
		return response.ErrorMessage(ctx, "Incorrect username/password")
	}

	//ctx.JSON(user)

	return response.SuccessMessage(ctx, "OK")
}
