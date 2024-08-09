# Use the official Golang image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go app for Linux
RUN go build -o main .

# Set executable permissions (if necessary)
RUN chmod +x /app/main

# Verify that the binary exists and is executable
RUN ls -l /app/main

# Set the entry point to the binary
ENTRYPOINT ["/app/main"]
