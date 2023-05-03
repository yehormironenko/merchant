package user

import (
	"github.com/rs/zerolog"
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
	us.logger.Info().Msg("service:RegisterNewUser")
	err := us.userRepo.RegisterUser(ctx, req)
	if err != nil {
		us.logger.Error().AnErr("error", err).Msg("user was not registered,")
		return err
	}
	us.logger.Info().Msg("the user has been successfully registered")
	return nil
}
