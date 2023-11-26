package requests

import "github.com/rs/zerolog"

type AuthUser struct {
	Username string `json:"username" validate:"required,gte=3"`
	Password string `json:"password" validate:"required,min=5"`
}

func (r AuthUser) MarshalZerologObject(e *zerolog.Event) {
	e.Str("Username", r.Username).
		Str("Password", r.Password)
}
