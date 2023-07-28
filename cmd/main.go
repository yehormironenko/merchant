package main

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"

	"merchant/config/client"
	"merchant/internal/repository"
	"merchant/internal/service/user"

	"merchant/config"
	"merchant/config/logger"
	"merchant/internal/controllers"
)

func main() {
	logger := logger.InitLogger()

	cfg, err := loadConfig()
	if err != nil {
		logger.Error().Msg("Cannot read config file: " + err.Error())
		return
	}

	engine := gin.Default()

	dynamoClient := client.NewDynamoClient(cfg.Dynamo)
	userRepo := repository.New(dynamoClient, cfg.Dynamo, logger)
	userService := user.New(userRepo, logger)

	controllers.Handlers(engine, userService, logger)

	engine.Run(cfg.Server.Port)
}

func loadConfig() (config.Config, error) {
	k := koanf.New(".")
	if err := k.Load(rawbytes.Provider([]byte(config.Yaml)), yaml.Parser()); err != nil {
		return config.Config{}, err
	}

	return config.Bind(k), nil
}
