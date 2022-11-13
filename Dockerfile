#base image
FROM golang:1.18-alpine AS build

# Set destination for COPY
RUN mkdir /go/src/app
ADD ./ /go/src/app
WORKDIR /go/src/app

# Download Go modules
RUN ls
RUN apk update && apk add --no-cache git && apk update && apk add ca-certificates && apk add build-base && rm -rf /var/cache/apk/*
RUN go build -a -o /main cmd/main.go

RUN go mod download


# Copy local code to the container image.

# Set env variable
#ENV AWS_ACCESS_KEY_ID=1111
#ENV AWS_REGION=eu-west-2
#ENV AWS_SECRET_ACCESS_KEY=1111

# Build the binary.

EXPOSE 8080

CMD [ "/main" ]
