package request

import (
	"github.com/RafikFarhad/hoax/utils"
	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Username string `validate:"required,min=5,max=30"example:"user_123"`
	Password string `validate:"required,min=1,max=50"example:"strong_password"`
}

func ValidateLoginRequest(req *LoginRequest) *validator.ValidationErrors {
	validate := utils.GetValidator()
	err := validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return &errors
	}
	return nil
}
