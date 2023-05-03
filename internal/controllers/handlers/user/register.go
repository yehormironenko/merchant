package user

import (
	"merchant/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/controllers/requests"
)

func Register(userService service.UserService, logger *zerolog.Logger) gin.HandlerFunc {

	return func(context *gin.Context) {
		var req requests.RegisterUser
		logger.Info().Msg("handlers:RegisterUserExecutor")
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			logger.Err(err).Msg("Bad request to register endpoint")
			return
		}
		err := userService.RegisterUser(context, req)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			logger.Err(err).Msg("Bad request to register endpoint")
			return
		}
		logger.Info().Object("data", req).Msg("user has been registered with following")

		context.JSON(http.StatusCreated, "Created")
	}
}
