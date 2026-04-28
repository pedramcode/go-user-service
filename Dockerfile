# Stage 1: Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build static binary with embedded certificates
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o user-service ./cmd/api

# Stage 2: Final Stage - SCRATCH (completely empty, no apk needed)
FROM scratch

# Copy only the binary and migrations
COPY --from=builder /build/user-service /user-service
COPY --from=builder /build/migrations /migrations

# Note: No ca-certificates, no tzdata, no shell

EXPOSE 8080

ENTRYPOINT ["/user-service"]