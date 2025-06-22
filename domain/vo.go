package domain

import (
	"errors"
	"regexp"
)

const EmailPattern = `^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`

var ErrInvalidEmail = errors.New("invalid email format")

type Email struct {
	address string
}

func isValidEmail(email string) bool {
	return regexp.MustCompile(EmailPattern).MatchString(email)
}

func NewEmail(address string) (Email, error) {
	if !isValidEmail(address) {
		return Email{}, ErrInvalidEmail
	}

	return Email{address: address}, nil
}
