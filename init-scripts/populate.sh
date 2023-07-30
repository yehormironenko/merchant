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


aws --endpoint-url=http://10.11.0.3:8000 dynamodb put-item  --table-name users  --item "{\"username\": {\"S\": \"user1\"}, \"fullname\": {\"S\": \"Mr. Bean\"}}"

aws --endpoint-url=http://10.11.0.3:8000 dynamodb put-item  --table-name users  --item "{\"username\": {\"S\": \"user2\"}, \"fullname\": {\"S\": \"John Smith\"}}"

aws --endpoint-url=http://10.11.0.3:8000 dynamodb put-item  --table-name users  --item "{\"username\": {\"S\": \"user3\"}, \"fullname\": {\"S\": \"Jan Kowalski\"}}"

echo "########### Selecting all data from a table ###########"
aws dynamodb scan --endpoint-url=http://10.11.0.3:8000 --table-name users