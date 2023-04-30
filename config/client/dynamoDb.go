package client

import (
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"golang.org/x/net/context"

	merchantConfig "merchant/config"
)

func NewDynamoClient(dynamoConf merchantConfig.Dynamo) *dynamodb.Client {
	httpClient := configureHttpClient(dynamoConf.HttpClient)
	ctx := context.Background()
	log.Info().Object("dynamoConfig", dynamoConf)
	opts := []func(*config.LoadOptions) error{
		config.WithHTTPClient(httpClient),
		config.WithRegion(dynamoConf.Region),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				if dynamoConf.Region == "local" {
					return aws.Endpoint{URL: dynamoConf.Url, SigningRegion: region}, nil
				}
				// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			},
		)),
	}

	conf, err := config.LoadDefaultConfig(ctx, opts...)

	if err != nil {
		log.Panic().Msg("failed to configure NewDynamoClient client")
	}
	return dynamodb.NewFromConfig(conf)
}

// configureHttpClient configures the 'resty' retryable client
func configureHttpClient(config merchantConfig.DynamoDbHttpClientConfig) *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConnsPerHost = config.Connections.MaxIdleConnectionPerHost
	transport.MaxConnsPerHost = config.Connections.MaxConnectionsPerHost
	transport.MaxIdleConns = config.Connections.MaxIdle
	return &http.Client{
		Transport: transport,
	}
}
