package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/controllers/requests"
	"merchant/internal/controllers/validators"
)

const AuthRequestCtxKey = "AuthRequestCtxKey"

func GetAuthUserRequestValidator(validator validators.Validators, logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var authRequest requests.AuthUser

		if err := c.ShouldBindJSON(&authRequest); err != nil {
			logger.Error().AnErr("invalid json request", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validator.Validate.Struct(authRequest); err != nil {
			logger.Error().AnErr("error validation structure", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Set(AuthRequestCtxKey, authRequest)

		c.Next()
	}
}
