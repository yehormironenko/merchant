package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"merchant/internal/service"
)

func Echo() gin.HandlerFunc {
	return func(context *gin.Context) {
		//TODO logs context
		//ctx := context.Request.Context()
		log.Println("Request to echo service")
		message := service.Echo()
		context.JSON(http.StatusOK, message)
	}
}
