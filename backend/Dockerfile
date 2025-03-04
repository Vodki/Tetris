# ========== DEVELOPMENT ========= 
FROM golang:1.24 AS development

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Install air for hot-reloading
RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]

# ========== PRODUCTION =========

# Stage 1: Build the application
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy Go module files first to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app ./server

# Stage 2: Create the production image
FROM alpine:latest

# Install SSL certificates (needed for HTTPS requests)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose the port your application runs on
EXPOSE 8080

# Command to run the application
CMD ["./app"]