package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/controllers/handlers"
	"merchant/internal/controllers/handlers/user"
	"merchant/internal/route"
)

func Handlers(engine *gin.Engine, logger *zerolog.Logger) {

	engine.GET(route.Echo, handlers.Echo(logger))

	public := engine.Group("/api")
	public.POST(route.Register, user.Register(logger))

}
