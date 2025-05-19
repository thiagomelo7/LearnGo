package domain

import validation "github.com/go-ozzo/ozzo-validation"

type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (u *Usuario) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Nome, validation.Required),
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Senha, validation.Required),
	)
}
