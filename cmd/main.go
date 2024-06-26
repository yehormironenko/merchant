package main

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/yehormironenko/reseller/pkg/api"

	"merchant/config/client"
	"merchant/internal/controllers/validators"
	"merchant/internal/repository"
	"merchant/internal/service/reseller"
	"merchant/internal/service/user"

	"merchant/config"
	"merchant/config/logger"
	"merchant/internal/controllers"
)

func main() {
	logger := logger.InitLogger()
	validator := validators.New(logger)
	cfg, err := loadConfig()
	if err != nil {
		logger.Error().Msg("Cannot read config file: " + err.Error())
		return
	}

	//externalClient
	resellerClient := api.NewResellerApiClient(cfg.Client.Endpoint)

	engine := gin.Default()

	dynamoClient := client.NewDynamoClient(cfg.Dynamo)
	userRepo := repository.New(dynamoClient, cfg.Dynamo, logger)
	userRegisterService := user.NewRegisterService(userRepo, logger)
	userAuthService := user.NewAuthService(userRepo, logger)

	resellerSearchBookService := reseller.NewSearchBookService(resellerClient, logger)

	controllers.Handlers(engine, userRegisterService, userAuthService, resellerSearchBookService, validator, logger)

	engine.Run(cfg.Server.Port)
}

func loadConfig() (config.Config, error) {
	k := koanf.New(".")
	if err := k.Load(rawbytes.Provider([]byte(config.Yaml)), yaml.Parser()); err != nil {
		return config.Config{}, err
	}

	return config.Bind(k), nil
}
