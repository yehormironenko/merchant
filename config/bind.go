package config

import (
	_ "embed"
	"log"

	"github.com/knadh/koanf"
)

//go:embed app-config.yaml
var Yaml string

type Config struct {
	Dynamo Dynamo
	Server Server
	Client Reseller
}

func Bind(kfg *koanf.Koanf) Config {
	appConfig := Config{}
	for path, field := range map[string]interface{}{
		"app.aws.dynamo":      &appConfig.Dynamo,
		"app.server":          &appConfig.Server,
		"app.client.reseller": &appConfig.Client,
	} {
		if err := kfg.Unmarshal(path, field); err != nil {
			log.Panic("config binding failed")
		}
	}

	if err := appConfig.Dynamo.Validate(); err != nil {
		panic(err)
	}

	return appConfig
}
