package main

import (
	"github.com/gin-gonic/gin"
	"merchant/middlewares"
	"merchant/pkg/controllers"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/login", controllers.Login)
	public.POST("/register", controllers.Register)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
