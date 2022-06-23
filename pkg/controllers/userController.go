package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	c "merchant/config"
	"merchant/database"
	"merchant/pkg/models"
	"net/http"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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

	token, err := LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	conf, _ := c.LoadConfig("./config/db", "config")
	svc := database.CreateSession(conf.DynamoDB.Endpoint)

	var user models.User
	var err error
	user, err = database.GetUser(username, svc)

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil {
		return "", err
	}

	return "successful login", nil
}
