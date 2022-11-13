#!/bin/bash
echo "########### Creating profile ###########"
aws configure set aws_access_key_id localstack --profile=default
aws configure set aws_secret_access_key localstack --profile=default
aws configure set default.region eu-west-2 --profile=default

echo "########### Listing profile ###########"
aws configure list
