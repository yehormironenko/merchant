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
	// TODO ctx for tracing in the future
	ur.logger.Info().Msg("registering new user")

	_, err := ur.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(ur.tables.Users),
		Item: map[string]types.AttributeValue{
			"username": &types.AttributeValueMemberS{Value: request.Username},
			"longname": &types.AttributeValueMemberS{Value: fmt.Sprintf("%s %s", request.Firstname, request.Surname)},
		},
	})
	//ur.logger.Log().Object("item", item) //TODO implement me
	if err != nil {
		return err
	}
	ur.logger.Info().Msg("New item has benn saved to dynamoDb")
	return nil
}

/*
func GetUser(username string) (user models.User, err error) {
	expr, _ := expression.NewBuilder().Build()

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String("users"),
	}
	result, err := svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	for _, i := range result.Items {
		item := models.User{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}

		if item.Username == username {
			return item, err
		}

	}
	log.Print("User not found...")

	return user, err
}
*/
