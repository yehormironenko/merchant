#!/bin/bash
echo "########### Creating profile ###########"
aws configure set aws_access_key_id 1 --profile=default
aws configure set aws_secret_access_key 1 --profile=default
aws configure set default.region local --profile=default

echo "########### Listing profile ###########"
aws configure list
