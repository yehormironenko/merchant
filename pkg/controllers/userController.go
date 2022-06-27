package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"merchant/database"
	"merchant/pkg/models"
	"merchant/utils"
	"net/http"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Longname string `json:"longname" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput
	log.Printf("New Input data %s ", input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password
	u.Longname = input.Longname

	err := database.SaveUser(u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := loginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func loginCheck(username string, password string) (string, error) {

	var user models.User
	var err error
	user, err = database.GetUser(username)

	if err != nil {
		return "", err
	}

	err = utils.VerifyPassword(password, user.Password)

	if err != nil {
		return "", err
	}

	return "successful login", nil
}
