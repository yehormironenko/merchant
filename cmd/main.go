package main

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/rs/zerolog/log"
	"merchant/config/client"
	"merchant/internal/repository"
	"merchant/internal/service/user"

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

	dynamoClient := client.NewDynamoClient(cfg.Dynamo)
	// repository
	userRepo := repository.New(dynamoClient, cfg.Dynamo, logger)

	// services
	userService := user.New(userRepo, logger)

	controllers.Handlers(engine, userService, logger)

	//	public.POST(route.Username, controllers.Username)
	//	protected := engine.Group("/api/admin")
	//protected.Use(middlewares.JwtAuthMiddleware())
	//protected.GET("/user", controllers.CurrentUser)

	engine.Run(cfg.Server.Port)
}
