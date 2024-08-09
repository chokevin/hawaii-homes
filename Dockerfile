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

# Set the entry point to the binary
ENTRYPOINT ["/app/main"]
