FROM golang:1.24 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy entire backend structure
COPY . .

# Build specifically the server package
RUN CGO_ENABLED=0 GOOS=linux go build -o tetris-backend ./server

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/tetris-backend .
EXPOSE 8080
CMD ["./tetris-backend"]