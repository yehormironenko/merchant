package repository

import (
	"context"

	"merchant/internal/controllers/requests"
)

type UserRepo interface {
	RegisterUser(ctx context.Context, request requests.RegisterUser) error
	GetUser(ctx context.Context, request requests.AuthUser) (*string, error)
}
