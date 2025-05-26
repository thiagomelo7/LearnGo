package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Usuario struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *Usuario) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Email, validation.Required),
	)
}
