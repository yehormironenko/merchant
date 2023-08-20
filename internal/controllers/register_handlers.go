package controllers

import (
	"net/http"

	middleware "merchant/internal/controllers/middlewares/user"
	"merchant/internal/controllers/validators"
	"merchant/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/controllers/handlers"
	"merchant/internal/controllers/handlers/user"
	"merchant/internal/route"
)

func Handlers(engine *gin.Engine, userService service.UserService, validator validators.Validators, logger *zerolog.Logger) {

	engine.GET(route.Echo, handlers.Echo(logger))

	public := engine.Group("/api")
	public.POST(route.Register, middleware.GetRegisterUserRequestvalidator(validator, logger), user.Register(userService, logger))

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "not found")
	})
}
