package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/service"
)

func Echo(logger *zerolog.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Info().Caller().Msg("handlers:EchoExecutor")
		message := service.Echo()
		context.JSON(http.StatusOK, message)
	}
}
