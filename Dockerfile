#FROM ubuntu:latest
#LABEL authors="Saurabh"
#
#ENTRYPOINT ["top", "-b"]

# Use the official Golang image as a base image
FROM golang:1.19

# Set the Current Working Directory inside the container
WORKDIR /GO-03

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

WORKDIR /GO-03/cmd

# Build the Go app
RUN go build -o /GO-03/main .

# Set the working directory back to /app
WORKDIR /GO-03

# Command to run the executable
CMD ["./main"]