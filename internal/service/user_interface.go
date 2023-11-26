package service

import (
	"golang.org/x/net/context"

	"merchant/internal/controllers/requests"
)

type UserService interface {
	RegisterUserService(ctx context.Context, req requests.RegisterUser) error
}

type AuthService interface {
	AuthUserService(ctx context.Context, req requests.AuthUser) (string, error)
}
