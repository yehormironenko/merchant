package user

import (
	"log"
	"merchant/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/controllers/requests"
)

func Register(userService service.UserService, logger *zerolog.Logger) gin.HandlerFunc {

	return func(context *gin.Context) {
		var req requests.RegisterUser
		logger.Info().Caller().Msg("handlers:RegisterUserExecutor")
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
		log.Printf("New Input data %s ", &req)

		context.JSON(http.StatusCreated, "Created")
	}
}
