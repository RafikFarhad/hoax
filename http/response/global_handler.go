package response

import (
	"github.com/RafikFarhad/hoax/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Something went wrong"
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	} else if len(err.Error()) > 0 {
		message = err.Error()
	}
	if needsJsonResponse(c) {
		return WithMessageAndCode(c, message, -1, code)
	} else {
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return c.Status(code).SendString(err.Error())
	}
}

type ApiResponse struct {
	ErrorCode int8        `json:"error"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func (base ApiResponse) Send(ctx *fiber.Ctx, extras ...int) error {
	httpStatus := fiber.StatusBadRequest
	if len(extras) > 0 {
		httpStatus = extras[0]
	}
	return ctx.Status(httpStatus).JSON(base)
}

func ForJson(ctx *fiber.Ctx, message string, data interface{}, errorCode int8, status int) error {
	return ApiResponse{
		ErrorCode: errorCode,
		Message:   message,
		Data:      data,
	}.Send(ctx, status)
}

func WithMessageAndCode(ctx *fiber.Ctx, message string, errorCode int8, status int) error {
	return ForJson(ctx, message, nil, errorCode, status)
}

func WithSuccess(ctx *fiber.Ctx, message string, data interface{}) error {
	return ForJson(ctx, message, data, 0, fiber.StatusOK)
}

func WithError(ctx *fiber.Ctx, message string, extras ...int) error {
	var errorCode int8 = 1
	httpStatus := fiber.StatusBadRequest
	if len(extras) > 0 {
		errorCode = int8(extras[0])
	}
	if len(extras) > 1 {
		httpStatus = extras[1]
	}
	return ForJson(ctx, message, nil, errorCode, httpStatus)
}

func WithValidationError(ctx *fiber.Ctx, fields *validator.ValidationErrors) error {
	responseData := map[string]string{}
	translator := utils.GetTranslator()
	for _, err := range *fields {
		responseData[err.Field()] = err.Translate(*translator)
	}
	return ForJson(ctx, "Validation error", responseData, 1, 422)
}

func needsJsonResponse(ctx *fiber.Ctx) bool {
	return strings.HasPrefix(
		string(ctx.Response().Header.ContentType()),
		"application/json",
	)
}
