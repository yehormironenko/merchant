package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"merchant/internal/service"
	"net/http"
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
