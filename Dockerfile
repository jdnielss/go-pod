# Use the official Golang image as a base image
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Install necessary dependencies for CGO
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Enable CGO and build the Go app
ENV CGO_ENABLED=0
RUN go build -o main .

# Use a minimal image for the final container
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
