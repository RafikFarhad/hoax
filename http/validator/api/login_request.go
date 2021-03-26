package api_validator

import (
	hoax_validator "github.com/RafikFarhad/hoax/http/validator"
	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Username string `validate:"required,min=5,max=30"`
	Password string `validate:"required,min=1,max=50"`
}

func ValidateLoginRequest(req *LoginRequest) *validator.ValidationErrors {
	validate := hoax_validator.GetValidator()
	err := validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return &errors
	}
	return nil
}
