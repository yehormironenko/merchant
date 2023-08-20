package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/controllers/requests"
	"merchant/internal/controllers/validators"
)

func GetRegisterUserRequestvalidator(validator validators.Validators, logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var registerRequest requests.RegisterUser

		if err := c.ShouldBindJSON(&registerRequest); err != nil {
			logger.Error().AnErr("invalid json request", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validator.Validate.Struct(registerRequest); err != nil {
			logger.Error().AnErr("error validation structure", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Next()
	}
}
