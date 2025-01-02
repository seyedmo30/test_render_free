# Stage 1: Build the application
FROM golang:1.22 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the module dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN GOARCH=amd64 GOOS=linux go build -o main .

# Stage 2: Create the runtime environment
FROM alpine:latest

RUN apk --no-cache add libc6-compat


# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]