package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"strings"
	"time"
)

type Dynamo struct {
	Region     string                   `koanf:"region"`
	Url        string                   `koanf:"endpoint"`
	Tables     Tables                   `koanf:"tables"`
	HttpClient DynamoDbHttpClientConfig `koanf:"httpClient"`
}

type Tables struct {
	Users string `koanf:"users"`
}

type DynamoDbHttpClientConfig struct {
	ConnectionTimeout time.Duration `koanf:"connectionTimeout"`
	Connections       Connections   `koanf:"connections"`
}

type Connections struct {
	MaxIdle                  int `koanf:"maxIdle"`
	MaxConnectionsPerHost    int `koanf:"maxConnPerHost"`
	MaxIdleConnectionPerHost int `koanf:"maxIdlePerHost"`
}

func (d Dynamo) MarshalZerologObject(e *zerolog.Event) {
	//TODO implement me
}

func (d Dynamo) Validate() error {
	var missing []string
	if d.Tables.Users == "" {
		missing = append(missing, "users")
	}
	if d.Region == "" {
		missing = append(missing, "region")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing DynamoDB config: %s", strings.Join(missing, ", "))
	}
	return nil
}
