#base image
FROM golang:1.18-alpine

# Set destination for COPY
WORKDIR /go/src/app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Set env variable
ENV AWS_ACCESS_KEY_ID=1111
ENV AWS_REGION=eu-west-2
ENV AWS_SECRET_ACCESS_KEY=1111

# Build the binary.
RUN go build -o /go/src/app

EXPOSE 8080

CMD [ "./main.exe" ]