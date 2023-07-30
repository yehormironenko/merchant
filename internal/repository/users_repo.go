package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/rs/zerolog"
	"golang.org/x/net/context"

	"merchant/config"
	"merchant/internal/controllers/requests"
)

const (
	username    = "username"
	fullname    = "fullname"
	email       = "email"
	phoneNumber = "phoneNumber"
)

type userRepo struct {
	db     *dynamodb.Client
	tables config.Tables
	logger *zerolog.Logger
}

func New(dynamoClient *dynamodb.Client, cfg config.Dynamo, logger *zerolog.Logger) UserRepo {
	return &userRepo{
		db:     dynamoClient,
		tables: cfg.Tables,
		logger: logger,
	}
}

// RegisterUser create new user in database if user doesn't exist
func (ur *userRepo) RegisterUser(ctx context.Context, request requests.RegisterUser) error {
	ur.logger.Info().Object("with data", request).Msg("registering new user object")
	condition := aws.String("attribute_not_exists(username) OR attribute_not_exists(email)") // TODO change the table

	resp, err := ur.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(ur.tables.Users),
		Item: map[string]types.AttributeValue{
			username:    &types.AttributeValueMemberS{Value: request.Username},
			fullname:    &types.AttributeValueMemberS{Value: fmt.Sprintf("%s %s", request.Firstname, request.Surname)},
			email:       &types.AttributeValueMemberS{Value: request.Email},
			phoneNumber: &types.AttributeValueMemberS{Value: request.PhoneNumber},
		},
		ConditionExpression: condition,
	})
	fmt.Println(resp)
	if err != nil {
		ur.logger.Error().AnErr("error", err).Msg("register operation ended with an error")
		return err
	}
	ur.logger.Info().Msg("new item has benn saved to dynamoDb")
	return nil
}
