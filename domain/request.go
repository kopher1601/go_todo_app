package domain

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type MemberRegisterRequest struct {
	Email    string `validate:"required,email"`
	Nickname string `validate:"required,min=3,max=20"`
	Password string `validate:"required,min=8,max=100"`
}

func (m *MemberRegisterRequest) Validate() error {
	return validate.Struct(m)
}
