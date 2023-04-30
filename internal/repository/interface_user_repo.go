package repository

import (
	"golang.org/x/net/context"
	"merchant/internal/controllers/requests"
)

type UserRepo interface {
	RegisterUser(ctx context.Context, request requests.RegisterUser) error
}
