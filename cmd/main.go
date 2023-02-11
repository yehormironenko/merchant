package main

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/rs/zerolog/log"

	"merchant/config"
	"merchant/config/logger"
	"merchant/internal/controllers"
)

func main() {

	logger := logger.InitLogger()
	k := koanf.New(".")
	if err := k.Load(rawbytes.Provider([]byte(config.Yaml)), yaml.Parser()); err != nil {
		log.Panic().Msg("Cannot read config file")
		return
	}

	cfg := config.Bind(k)
	engine := gin.Default()

	controllers.Handlers(engine, logger)

	//	public.POST(route.Login, controllers.Login)
	//	protected := engine.Group("/api/admin")
	//protected.Use(middlewares.JwtAuthMiddleware())
	//protected.GET("/user", controllers.CurrentUser)

	engine.Run(cfg.Server.Port)
}
