# Stage 1: Build the Go binary
FROM golang:1.22.5 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/home/home.go

# Stage 2: Create a minimal image with the Go binary
FROM alpine:latest

# Install necessary CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/main /main

# Expose port 50051 to the outside world
EXPOSE 50051

# Set executable permissions (if not already set)
RUN chmod +x /main

# Command to run the executable
ENTRYPOINT ["/main"]