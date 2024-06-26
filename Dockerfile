# First stage: Build the Go application
FROM golang:latest AS builder

# Set the working directory
WORKDIR /bin

# Copy the Go application files to the container
COPY . .

# Install any Go dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o go-pod

# Second stage: Create the runtime environment
FROM gcr.io/distroless/base-debian11

# Set the working directory
WORKDIR /bin

# Copy the Go binary from the build stage
COPY --from=builder /bin/go-pod /bin/go-pod

# Define the entrypoint command
ENTRYPOINT ["/bin/go-pod"]
