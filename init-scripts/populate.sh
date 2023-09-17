#!/bin/bash

echo "########### Creating table with global secondary index ###########"

aws dynamodb --endpoint-url=http://10.11.0.3:8000 create-table \
  --table-name users \
  --attribute-definitions \
    AttributeName=userId,AttributeType=S \
    AttributeName=username,AttributeType=S \
    AttributeName=email,AttributeType=S \
  --key-schema \
    AttributeName=userId,KeyType=HASH \
	AttributeName=username,KeyType=RANGE \
  --provisioned-throughput \
    ReadCapacityUnits=10,WriteCapacityUnits=5 \
  --global-secondary-indexes \
    '[
        {
            "IndexName": "usernameIndex",
            "KeySchema": [{"AttributeName":"username","KeyType":"HASH"}],
            "Projection": {
                "ProjectionType": "INCLUDE",
                "NonKeyAttributes": ["username"]
            },
            "ProvisionedThroughput": {
                "ReadCapacityUnits": 10,
                "WriteCapacityUnits": 5
            }
        },
        {
            "IndexName": "emailIndex",
            "KeySchema": [{"AttributeName":"email","KeyType":"HASH"}],
            "Projection": {
                "ProjectionType": "INCLUDE",
                "NonKeyAttributes": ["userId"]
            },
            "ProvisionedThroughput": {
                "ReadCapacityUnits": 10,
                "WriteCapacityUnits": 5
            }
        }
    ]'

echo "################# Table created ###################"

# Create 5 test items with UUIDv4 userId values
aws --endpoint-url=http://10.11.0.3:8000 dynamodb put-item --table-name users --item '{"userId": {"S": "'0ebf0fdb-09cb-46c8-9eb9-12c83a540cd5'"}, "username": {"S": "user1"}, "fullname": {"S": "John Doe"}, "email": {"S": "user1@example.com"}, "phoneNumber": {"S": "+1234567890"}}'

aws --endpoint-url=http://10.11.0.3:8000 dynamodb put-item --table-name users --item '{"userId": {"S": "'d8632d18-6286-4a7c-8ff7-5679a65234f6'"}, "username": {"S": "user2"}, "fullname": {"S": "Jane Smith"}, "email": {"S": "user2@example.com"}}'

aws --endpoint-url=http://10.11.0.3:8000 dynamodb put-item --table-name users --item '{"userId": {"S": "'8bca68dd-f85e-429d-9c99-5e9d8c05dc82'"}, "username": {"S": "user3"}, "fullname": {"S": "Alice Johnson"}, "email": {"S": "user3@example.com"}}'

aws --endpoint-url=http://10.11.0.3:8000 dynamodb put-item --table-name users --item '{"userId": {"S": "'753fc301-4d01-41e2-9c42-73f01eb4cbfe'"}, "username": {"S": "user4"}, "fullname": {"S": "Bob Brown"}, "email": {"S": "user4@example.com"}, "phoneNumber": {"S": "+9876543210"}}'

aws --endpoint-url=http://10.11.0.3:8000 dynamodb put-item --table-name users --item '{"userId": {"S": "'ae129f55-a39b-4b55-b3aa-dbef3497da14'"}, "username": {"S": "user5"}, "fullname": {"S": "Eve Williams"}, "email": {"S": "user5@example.com"}}'

echo "########### Selecting all data from a table ###########"
aws dynamodb scan --endpoint-url=http://10.11.0.3:8000 --table-name users