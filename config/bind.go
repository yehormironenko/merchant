package config

import (
	_ "embed"
	"github.com/knadh/koanf"
	"log"
)

//go:embed app-config.yaml
var Yaml string

type Config struct {
	Dynamo Dynamo
	Server Server
}

func Bind(kfg *koanf.Koanf) Config {
	appConfig := Config{}
	for path, field := range map[string]interface{}{
		"app.aws":    &appConfig.Dynamo,
		"app.server": &appConfig.Server,
	} {
		if err := kfg.Unmarshal(path, field); err != nil {
			log.Panic("config binding failed")
		}
	}
	return appConfig
}
