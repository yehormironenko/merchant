package user

import (
	"net/http"

	"merchant/internal/controllers/middlewares/user"
	"merchant/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/controllers/requests"
)

func Register(userService service.UserService, logger *zerolog.Logger) gin.HandlerFunc {

	return func(context *gin.Context) {
		//	var req requests.RegisterUserService
		logger.Info().Msg("handlers:RegisterUserExecutor")
		req := context.Value(user.RegisterRequestCtxKey)

		userRequest, err := userService.RegisterUserService(context, req.(requests.RegisterUser))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			logger.Err(err).Msg("Bad request to register endpoint")
			return
		}
		logger.Info().Object("data", req.(requests.RegisterUser)).Msg("user has been registered with following")

		context.JSON(http.StatusCreated, userRequest)
	}
}
