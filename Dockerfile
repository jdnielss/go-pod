# Use the official Golang image as a base image
FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and build files
COPY . . 
RUN go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Use a minimal base image to reduce the image size
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Set the entry point for the container
ENTRYPOINT ["./app"]
