package controllers

import (
	"net/http"

	"merchant/internal/controllers/handlers/reseller"
	middleware "merchant/internal/controllers/middlewares/user"
	"merchant/internal/controllers/validators"
	"merchant/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/controllers/handlers"
	"merchant/internal/controllers/handlers/user"
	"merchant/internal/route"
)

func Handlers(engine *gin.Engine, userRegisterService service.UserService, userAuthService service.AuthService, searchBookService service.ResellerSearchService, validator validators.Validators, logger *zerolog.Logger) {

	engine.GET(route.Echo, handlers.Echo(logger))

	public := engine.Group("/api")
	public.POST(route.Register, middleware.GetRegisterUserRequestValidator(validator, logger), user.Register(userRegisterService, logger))
	public.POST(route.Auth, middleware.GetAuthUserRequestValidator(validator, logger), user.Auth(userAuthService, logger))

	public.POST(route.Search, reseller.SearchBook(searchBookService, logger))

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "not found")
	})
}
