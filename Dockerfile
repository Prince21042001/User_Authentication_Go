# Use a Golang base image
FROM golang:1.23.1-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a lightweight image for the final image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary and templates
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates

# Expose the port your application listens on
EXPOSE 8080

# Command to run the application
CMD ["./main"]