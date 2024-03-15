# Use the official Golang image
FROM golang:1.22.0 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o CC-Compiler-Go .

# Use a minimal base image to reduce size
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/CC-Compiler-Go .

# Copy the models and controllers folders from the root directory
COPY --from=builder /app/models /app/models
COPY --from=builder /app/controllers /app/controllers

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./CC-Compiler-Go"]