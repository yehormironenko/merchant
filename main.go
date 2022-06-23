package main

import (
	"github.com/gin-gonic/gin"
	"merchant/pkg/controllers"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/login", controllers.Login)

	r.Run(":8080")

}
