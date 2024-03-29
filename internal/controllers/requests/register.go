package requests

import "github.com/rs/zerolog"

type RegisterUser struct {
	Username    string  `json:"username" validate:"required,gte=3"`
	Password    *string `json:"password,omitempty" validate:"required,min=6"`
	Firstname   string  `json:"firstname" validate:"required,gte=3"`
	Surname     string  `json:"surname" validate:"required,gte=3"`
	Email       string  `json:"email" validate:"required,email"`
	PhoneNumber string  `json:"phoneNumber,omitempty" validate:"omitempty,e164"`
}

func (r RegisterUser) MarshalZerologObject(e *zerolog.Event) {
	e.Str("Username", r.Username).
		Str("Password", *r.Password).
		Str("Firstname", r.Firstname).
		Str("Surname", r.Surname).
		Str("Email", r.Email).
		Str("PhoneNumber", r.PhoneNumber)
}
