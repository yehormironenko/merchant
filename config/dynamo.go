package config

import "time"

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
