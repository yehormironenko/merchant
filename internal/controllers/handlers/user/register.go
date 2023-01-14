package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"merchant/internal/controllers/requests"
)

func Register() gin.HandlerFunc {

	return func(context *gin.Context) {
		var req requests.RegisterUser

		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Printf("New Input data %s ", &req)

		context.JSON(http.StatusOK, "Created")
	}
}
