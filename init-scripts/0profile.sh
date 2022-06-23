#!/bin/bash
echo "########### Creating profile ###########"
aws configure set aws_access_key_id 1111 --profile=default
aws configure set aws_secret_access_key 1111 --profile=default
aws configure set region eu-west-2 --profile=default

echo "########### Listing profile ###########"
aws configure list
