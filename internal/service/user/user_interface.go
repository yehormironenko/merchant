package user

import (
	"golang.org/x/net/context"
	"merchant/internal/controllers/requests"
)

type UserService interface {
	RegisterUser(ctx context.Context, req requests.RegisterUser)
}
