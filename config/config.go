package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	DynamoDB DatabaseConfigurations
}

//TODO this is model
// Config exported
type DatabaseConfigurations struct {
	DBName     string `mapstructure:"dbname"`
	Endpoint   string `mapstructure:"endpoint"`
	DBHost     string `mapstructure:"dbhost"`
	DBPort     string `mapstructure:"dbport"`
	DBUser     string `mapstructure:"dbuser"`
	DBPassword string `mapstructure:"dbpassword"`
}

func LoadConfig(path string, file string) (configuration Configurations, err error) {
	// Set the file name of the configurations file
	viper.SetConfigName(file)
	// Set the path to look for the configurations file
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	var conf Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return conf, err
}