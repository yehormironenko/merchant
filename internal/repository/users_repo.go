package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"golang.org/x/net/context"

	"merchant/config"
	"merchant/internal/controllers/requests"
)

const (
	userId       = "userId"
	username     = "username"
	passwordHash = "password hash"
	fullname     = "fullname"
	email        = "email"
	phoneNumber  = "phoneNumber"
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

func (ur *userRepo) RegisterUser(ctx context.Context, request requests.RegisterUser) error {
	ur.logger.Info().Object("with data", request).Msg("registering new user object")

	// Check if the email is unique
	emailExists, err := ur.checkAttributeExists(ctx, "email", request.Email)
	if err != nil {
		ur.logger.Error().AnErr("error", err).Msg("register operation ended with an error")
		return err
	}

	if emailExists {
		ur.logger.Error().Msg("email already exists")
		return fmt.Errorf("email already exists")
	}

	// Check if the username is unique
	usernameExists, err := ur.checkAttributeExists(ctx, "username", request.Username)
	if err != nil {
		ur.logger.Error().AnErr("error", err).Msg("register operation ended with an error")
		return err
	}

	if usernameExists {
		ur.logger.Error().Msg("username already exists")
		return fmt.Errorf("username already exists")
	}

	// Both email and username are unique, proceed with the PutItem operation
	_, err = ur.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(ur.tables.Users),
		Item: map[string]types.AttributeValue{
			userId:       &types.AttributeValueMemberS{Value: uuid.New().String()},
			username:     &types.AttributeValueMemberS{Value: request.Username},
			passwordHash: &types.AttributeValueMemberS{Value: request.Password},
			fullname:     &types.AttributeValueMemberS{Value: fmt.Sprintf("%s %s", request.Firstname, request.Surname)},
			email:        &types.AttributeValueMemberS{Value: request.Email},
			phoneNumber:  &types.AttributeValueMemberS{Value: request.PhoneNumber},
		},
	})

	if err != nil {
		ur.logger.Error().AnErr("error", err).Msg("register operation ended with an error")
		return err
	}

	ur.logger.Info().Msg("new item has been saved to DynamoDB")
	return nil
}

// checkAttributeExists checks if the given attribute already exists in the table
func (ur *userRepo) checkAttributeExists(ctx context.Context, attributeName, attributeValue string) (bool, error) {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(ur.tables.Users),
		IndexName: aws.String(attributeName + "Index"),
		KeyConditions: map[string]types.Condition{
			attributeName: {
				ComparisonOperator: types.ComparisonOperatorEq,
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberS{Value: attributeValue},
				},
			},
		},
		Select: types.SelectCount,
	}

	result, err := ur.db.Query(ctx, queryInput)
	if err != nil {
		return false, err
	}

	return result.Count > 0, nil
}
