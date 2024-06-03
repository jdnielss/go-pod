FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download
 
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
