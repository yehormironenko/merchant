#!/bin/bash


echo "########### Creating table with global secondary index ###########"

aws dynamodb --endpoint-url=http://localhost:4566 create-table \
               --table-name users \
               --attribute-definitions \
	       AttributeName=id,AttributeType=S \
               AttributeName=username,AttributeType=S \
               --key-schema \
               AttributeName=id,KeyType=HASH \
	       AttributeName=username,KeyType=RANGE \
	       --provisioned-throughput \
	       ReadCapacityUnits=10,WriteCapacityUnits=5

echo "################# Table created ###################"

password1='$2a$12$nhKBZ5c0N0lWii30ABlFte85OobnbPypFj4KungPtvQ4bmwxSXJfi' #`echo -n 1111 | sha256sum | cut -d " " -f1`
password2='$2a$12$VpQhPWGXHKt5pgdjM9Je4.QiZacXKdzXhuaR0bibLb5tnwkie4FFe' #`echo -n 2222 | sha256sum | cut -d " " -f1`
password3='$2a$12$R4axbZ.Ol78NIUe6Edk9Ju.c7KPUv8JT9g0ipkMV5GTySYpVfQyva' #`echo -n 3333 | sha256sum | cut -d " " -f1`

echo $passord3

aws --endpoint-url=http://localhost:4566 dynamodb put-item  --table-name users  --item "{\"id\": {\"S\": \"1\"}, \"username\": {\"S\": \"user1\"}, \"password\": {\"S\": \"$password1\"}, \"longName\": {\"S\": \"Mr. Bean\"}}"

aws --endpoint-url=http://localhost:4566 dynamodb put-item  --table-name users  --item "{\"id\": {\"S\": \"2\"}, \"username\": {\"S\": \"user2\"}, \"password\": {\"S\": \"$password2\"}, \"longName\": {\"S\": \"John Smith\"}}"

aws --endpoint-url=http://localhost:4566 dynamodb put-item  --table-name users  --item "{\"id\": {\"S\": \"3\"}, \"username\": {\"S\": \"user3\"}, \"password\": {\"S\": \"$password3\"}, \"longName\": {\"S\": \"Jan Kowalski\"}}"

echo "########### Selecting all data from a table ###########"
aws dynamodb scan --endpoint-url=http://localhost:4566 --table-name users
