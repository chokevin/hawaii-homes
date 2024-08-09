# Use the official Golang image as the base image
FROM golang:1.22.5

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Run go mod download
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go app for Linux
RUN go build -o main ./cmd/home/home.go

# Set executable permissions (if necessary)
RUN chmod +x /app/main

# Verify that the binary exists and is executable
RUN ls -l /app/main

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 50051

# Set the entry point to the binary
ENTRYPOINT ["/app/main"]
