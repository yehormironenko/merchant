package database

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	c "merchant/config"
	"merchant/internal/models"
)

var conf, _ = c.LoadConfig("./cmd/config/db", "config")
var svc = CreateSession(conf.DynamoDB.Endpoint)

func CreateSession(endpoint string) *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess, aws.NewConfig().WithEndpoint(endpoint))
	return svc
}

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

func SaveUser(user models.User) error {

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Fatalf("Got error marshalling new user: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("users"),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	log.Println("Successfully added item  to table " + "users")
	log.Printf("Added user: %s", av)

	return nil
}
