package response

import (
	validator_singleton "github.com/RafikFarhad/hoax/http/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessMessage(ctx *fiber.Ctx, message string) error {
	return SendApiResponse(ctx, 200, message, nil)
}

func SuccessMessageWithStatus(ctx *fiber.Ctx, message string, status int) error {
	return SendApiResponse(ctx, status, message, nil)
}

func SuccessMessageWithData(ctx *fiber.Ctx, data interface{}) error {
	return SendApiResponse(ctx, 200, "", data)
}

func ErrorMessage(ctx *fiber.Ctx, message string) error {
	return SendApiResponse(ctx, 400, message, nil)
}

func ErrorMessageWithStatus(ctx *fiber.Ctx, message string, status int) error {
	return SendApiResponse(ctx, status, message, nil)
}

func ValidationError(ctx *fiber.Ctx, fields *validator.ValidationErrors) error {
	responseData := map[string]string{}
	translator := validator_singleton.GetTranslator()
	for _, err := range *fields {
		responseData[err.Field()] = err.Translate(*translator)
	}
	return SendApiResponse(ctx, 422, "Validation error", responseData)
}

func SendApiResponse(ctx *fiber.Ctx, status int, message string, data interface{}) error {
	return ctx.Status(status).JSON(ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
