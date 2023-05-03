#base image
FROM golang:1.18-alpine AS build

# Set destination for COPY
RUN mkdir /go/src/app
ADD ./ /go/src/app
WORKDIR /go/src/app

# Download Go modules
RUN ls
RUN apk add --update git && apk add --update ca-certificates  && apk add --update curl && apk add --update unzip && apk add --update py-pip && apk add build-base && rm -rf /var/cache/apk/*
RUN go build -a -o /main cmd/main.go
RUN go mod download
# AWS CLI installation commands
#RUN	curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
#RUN	unzip awscliv2.zip
#RUN ./aws/install -i /usr/local/aws -b /usr/local/bin/aws
RUN python -m pip install awscli

EXPOSE 8080

CMD [ "/main" ]
