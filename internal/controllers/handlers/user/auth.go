package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/controllers/middlewares/user"
	"merchant/internal/controllers/requests"
	"merchant/internal/service"
)

func Auth(authService service.AuthService, logger *zerolog.Logger) gin.HandlerFunc {

	return func(context *gin.Context) {
		logger.Info().Msg("handlers:AuthenticationUserExecutor")
		req := context.Value(user.AuthRequestCtxKey)

		token, err := authService.AuthUserService(context, req.(requests.AuthUser))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			logger.Err(err).Msg("Bad request to authentication endpoint")
			return
		}
		logger.Info().Object("data", req.(requests.AuthUser)).Msg("user has been authenticated with following")

		context.JSON(http.StatusOK, token)
	}
}
