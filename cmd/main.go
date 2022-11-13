package main

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"merchant/config"
	"merchant/internal/controllers/handlers"
	"merchant/internal/controllers/handlers/user"
	"merchant/internal/route"
)

func main() {
	k := koanf.New(".")
	err := k.Load(rawbytes.Provider([]byte(config.Yaml)), yaml.Parser())

	if err != nil {
		return
	}

	cfg := config.Bind(k)
	r := gin.Default()

	r.GET(route.Echo, handlers.Echo())

	public := r.Group("/api")
	public.POST(route.Register, user.Register())

	//	public.POST(route.Login, controllers.Login)
	//	protected := r.Group("/api/admin")
	//protected.Use(middlewares.JwtAuthMiddleware())
	//protected.GET("/user", controllers.CurrentUser)

	r.Run(cfg.Server.Port)
}
