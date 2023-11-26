package user

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"

	"merchant/internal"
	"merchant/internal/controllers/requests"
	"merchant/internal/repository"
)

type AuthService struct {
	userRepo repository.UserRepo
	logger   *zerolog.Logger
}

func NewAuthService(userRepo repository.UserRepo, logger *zerolog.Logger) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (as AuthService) AuthUserService(ctx context.Context, req requests.AuthUser) (string, error) {
	as.logger.Info().Msg("service:AuthUser")

	username, err := as.userRepo.GetUser(ctx, req)
	if err != nil {
		as.logger.Error().AnErr("error", err).Msg("user was not authenticated")
		return "", err
	}

	if username == nil {
		as.logger.Error().AnErr("error", err).Msg("user not found")
		return "", fmt.Errorf("user not found")
	}
	token, err := generateJWT(*username)
	if err != nil {
		as.logger.Error().AnErr("error", err).Msg("cannot generate token")
		return "", err
	}

	as.logger.Info().Msg("returns auth token")
	return token, nil
}

func generateJWT(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(internal.TokenTTL).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user"] = username

	tokenString, err := token.SignedString([]byte(internal.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
