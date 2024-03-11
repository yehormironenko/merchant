package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"

	"merchant/config"
	"merchant/internal"
	"merchant/internal/controllers/requests"
)

const (
	userId       = "userId"
	username     = "username"
	passwordHash = "passwordHash"
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
			email:        &types.AttributeValueMemberS{Value: request.Email},
			passwordHash: &types.AttributeValueMemberS{Value: *request.Password},
			fullname:     &types.AttributeValueMemberS{Value: fmt.Sprintf("%s %s", request.Firstname, request.Surname)},
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

func (ur *userRepo) GetUser(ctx context.Context, request requests.AuthUser) (*string, error) { //TODO user to return model
	ur.logger.Info().Object("with data", request).Msg("authenticate user object")
	var storedPassword string
	// Define the input parameters for the GetItem operation
	user, err := ur.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(ur.tables.Users),
		Key: map[string]types.AttributeValue{
			username: &types.AttributeValueMemberS{Value: request.Username},
		},
	})

	if err != nil {
		ur.logger.Error().AnErr("error", err).Msg("authentication operation ended with an error")
		return nil, err
	}

	// Check if the user was found
	if user.Item == nil {
		ur.logger.Info().Msg(fmt.Sprintf("user not found: %s", request.Username))
		return nil, fmt.Errorf("user not found")
	}

	storedPasswordAttribute := user.Item[passwordHash]

	err = attributevalue.Unmarshal(storedPasswordAttribute, &storedPassword)
	if err != nil {
		ur.logger.Info().Msg(fmt.Sprintf("cannot unmarshal password: %s", storedPasswordAttribute))
		return nil, fmt.Errorf("service error")
	}

	if !verifyPassword(storedPassword, request.Password) {
		ur.logger.Info().Msg(fmt.Sprintf("incorect password"))
		return nil, fmt.Errorf("user not found")
	}

	ur.logger.Info().Msg(fmt.Sprintf("username '%s' exists", request.Username))
	return &request.Username, nil
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

// verifyPassword checks if a user-provided password matches the stored hashed password
func verifyPassword(storedHashedPassword string, userProvidedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(userProvidedPassword+internal.Salt))
	return err == nil
}
