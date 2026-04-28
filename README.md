# User Service

Golang user authentication, authorization and management service.

## Architecture

This project follows **Clean Architecture** (also known as Hexagonal Architecture) principles to maintain separation of concerns, testability, and independence from external frameworks.

## Dependencies

### Database

- Postgres
- Redis

## Development

### Tools

install:
* `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
* `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
* `go install github.com/swaggo/swag/cmd/swag@latest`
* `go install github.com/vektra/mockery/v2@latest`