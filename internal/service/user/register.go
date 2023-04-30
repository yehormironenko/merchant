package user

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"merchant/internal/controllers/requests"
	"merchant/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepo
	logger   *zerolog.Logger
}

func New(userRepo repository.UserRepo, logger *zerolog.Logger) *UserService {
	return &UserService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (us *UserService) RegisterUser(ctx context.Context, req requests.RegisterUser) error {
	log.Info().Msg("service:RegisterNewUser")
	err := us.userRepo.RegisterUser(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
