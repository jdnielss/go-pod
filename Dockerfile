FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download
 
RUN CGO_ENABLED=0 go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
