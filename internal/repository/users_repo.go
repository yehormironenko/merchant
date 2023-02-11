package repository

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/rs/zerolog"

	"merchant/config"
)

type userRepo struct {
	db     *dynamodb.Client
	tables config.Tables
	//paymentInstrumentIndexes      config.PaymentInstrumentIndexes
	//paymentInstrumentTokenIndexes config.PaymentInstrumentTokenIndexes
	logger *zerolog.Logger
}

func New(dynamoClient *dynamodb.Client, cfg config.Dynamo, logger *zerolog.Logger) UserRepo {
	return &userRepo{
		db:     dynamoClient,
		tables: cfg.Tables,
		//	paymentInstrumentIndexes:      cfg.PaymentInstrumentIndexes,
		//	paymentInstrumentTokenIndexes: cfg.PaymentInstrumentTokenIndexes,
		logger: logger,
	}
}
