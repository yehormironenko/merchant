package user

import (
	"context"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"

	"merchant/internal"
	"merchant/internal/controllers/requests"
	"merchant/internal/repository"
)

type RegisterService struct {
	userRepo repository.UserRepo
	logger   *zerolog.Logger
}

func NewRegisterService(userRepo repository.UserRepo, logger *zerolog.Logger) *RegisterService {
	return &RegisterService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (rs RegisterService) RegisterUserService(ctx context.Context, req requests.RegisterUser) (requests.RegisterUser, error) {
	rs.logger.Info().Msg("service:RegisterNewUser")
	req.Password = rs.generatePasswordHash(req.Password)
	err := rs.userRepo.RegisterUser(ctx, req)
	if err != nil {
		rs.logger.Error().AnErr("error", err).Msg("user was not registered,")
		return requests.RegisterUser{}, err
	}
	rs.logger.Info().Msg("the user has been successfully registered")
	req.Password = nil
	return req, nil
}

func (rs RegisterService) generatePasswordHash(password *string) *string {
	// Generate a salted hash of the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password+internal.Salt), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	stringHashedPassword := string(hashedPassword)
	return &stringHashedPassword
}
