# Use Go 1.21 as the base image for building
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go module files
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/my-app

# Create a minimal image for running the application
FROM gcr.io/distroless/static-debian12

# Copy the binary from the builder stage
COPY --from=builder /app/my-app /my-app

# Set the entry point for the container
ENTRYPOINT ["/my-app"]
