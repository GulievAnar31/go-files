package errors

import (
	"errors"

	"github.com/fatih/color"
)

var (
	ErrInvalidLogin = errors.New("Invalid Login")
	ErrInvalidURL   = errors.New("Invalid Url")
)

func GetErrorString(err error) {
	switch {
	case errors.Is(err, ErrInvalidLogin):
		color.Red("Логин должен быть больше 3 символов")
	case errors.Is(err, ErrInvalidURL):
		color.Red("Неверный формат URL")
	}
}
