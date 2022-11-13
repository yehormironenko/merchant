#!/bin/bash

echo "########### Creating table with global secondary index ###########"

aws dynamodb --endpoint-url=http://localhost:4566 create-table \
               --table-name users \
               --attribute-definitions \
	             AttributeName=longname,AttributeType=S \
               AttributeName=username,AttributeType=S \
               --key-schema \
               AttributeName=longname,KeyType=HASH \
              AttributeName=username,KeyType=RANGE \
	            --provisioned-throughput \
	            ReadCapacityUnits=10,WriteCapacityUnits=5

echo "################# Table created ###################"


aws --endpoint-url=http://localhost:4566 dynamodb put-item  --table-name users  --item "{\"username\": {\"S\": \"user1\"}, \"longname\": {\"S\": \"Mr. Bean\"}}"

aws --endpoint-url=http://localhost:4566 dynamodb put-item  --table-name users  --item "{\"username\": {\"S\": \"user2\"}, \"longname\": {\"S\": \"John Smith\"}}"

aws --endpoint-url=http://localhost:4566 dynamodb put-item  --table-name users  --item "{\"username\": {\"S\": \"user3\"}, \"longname\": {\"S\": \"Jan Kowalski\"}}"



echo "########### Selecting all data from a table ###########"
aws dynamodb scan --endpoint-url=http://localhost:4566 --table-name users
