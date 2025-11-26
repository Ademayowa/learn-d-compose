FROM golang:1.23-alpine AS builder
WORKDIR /app

# Set GOTOOLCHAIN to allow using newer Go version features
ENV GOTOOLCHAIN=auto

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main ./cmd/main.go

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy binary only
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]