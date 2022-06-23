package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"log"
	"merchant/pkg/models"
)

/*func ConnectToDatabase() {

	conf, _ := c.LoadConfig("./config/db", "config")
	// svc := createSession(conf.DynamoDB.Endpoint)
	log.Print("Connecting to database ", conf.DynamoDB.Endpoint)

}*/

func CreateSession(endpoint string) *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess, aws.NewConfig().WithEndpoint(endpoint))
	return svc
}

func GetUser(username string, svc *dynamodb.DynamoDB) (user models.User, err error) {
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
