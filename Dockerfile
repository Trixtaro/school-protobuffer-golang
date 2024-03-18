# Use the official Golang image as the base image
FROM golang:1.22

# Install the protobuf compiler
RUN apt-get update && apt-get install -y protobuf-compiler

# Set the working directory inside the container
WORKDIR /app

# Copy only the go.mod and go.sum files to leverage Docker's layer caching
COPY go.mod go.sum ./

# Install the protoc-gen-go plugin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

# Add protoc-gen-go to the PATH
RUN export GO_PATH=~/go
RUN export PATH=$PATH:/$GO_PATH/bin

# Download the dependencies
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Run the make file commands
RUN make compile

# Build the Golang app
RUN go build -o bin/start cmd/main.go

# Expose the port that the app runs on
EXPOSE 5060
