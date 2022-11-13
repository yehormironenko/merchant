package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"merchant/database"
	"merchant/internal/controllers/requests"
	"merchant/internal/models"
	"merchant/utils/token"
	"net/http"
)

func Register(c *gin.Context) {

	var input requests.RegisterUser
	log.Printf("New Input data %s ", &input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Longname = input.Longname

	err := database.SaveUser(u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

func CurrentUser(c *gin.Context) {

	username, err, expDate := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := database.GetUser(username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u, "expiration date": expDate})
}
