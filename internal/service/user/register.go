package user

import (
	"crypto/sha1"
	"fmt"

	"github.com/rs/zerolog"
	"golang.org/x/net/context"

	"merchant/internal/controllers/requests"
	"merchant/internal/repository"
)

const salt = "Gvug2HK4HDNCXSfW3Fsw2RmOl"

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

func (u UserService) RegisterUser(ctx context.Context, req requests.RegisterUser) error {
	u.logger.Info().Msg("service:RegisterNewUser")
	req.Password = u.generatePasswordHash(req.Password)
	err := u.userRepo.RegisterUser(ctx, req)
	if err != nil {
		u.logger.Error().AnErr("error", err).Msg("user was not registered,")
		return err
	}
	u.logger.Info().Msg("the user has been successfully registered")
	return nil
}

func (u UserService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
