package response

import (
	"github.com/RafikFarhad/hoax/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessMessage(ctx *fiber.Ctx, message string) error {
	return SendApiResponse(ctx, 200, 200, message, nil)
}

func SuccessMessageWithStatus(ctx *fiber.Ctx, message string, code int) error {
	return SendApiResponse(ctx, code, code, message, nil)
}

func SuccessMessageWithData(ctx *fiber.Ctx, data interface{}) error {
	return SendApiResponse(ctx, 200, 200, "", data)
}

func ErrorMessage(ctx *fiber.Ctx, message string) error {
	return SendApiResponse(ctx, 400, 400, message, nil)
}

func ErrorMessageWithStatus(ctx *fiber.Ctx, message string, code int) error {
	return SendApiResponse(ctx, code, code, message, nil)
}

func ValidationError(ctx *fiber.Ctx, fields *validator.ValidationErrors) error {
	responseData := map[string]string{}
	translator := utils.GetTranslator()
	for _, err := range *fields {
		responseData[err.Field()] = err.Translate(*translator)
	}
	return SendApiResponse(ctx, 422, 422, "Validation error", responseData)
}

func SendApiResponse(ctx *fiber.Ctx, code int, httpStatus int, message string, data interface{}) error {
	return ctx.Status(httpStatus).JSON(ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
