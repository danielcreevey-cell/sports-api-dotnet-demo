# Sports API - Go Port

This is a Go port of the .NET Clean Architecture Todo API. It follows Clean Architecture principles with clearly separated layers.

## Status

**Task 1 of migration**: Domain models, DTOs, and project scaffolding.

## Project Structure

```
go-api/
├── cmd/
│   └── api/
│       └── main.go              # Entry point (HTTP server with health check)
├── internal/
│   ├── domain/                  # Domain layer: entities, value objects, events
│   ├── application/
│   │   ├── dto/                 # Data transfer objects
│   │   └── port/                # Repository interfaces (Task 2)
│   ├── infrastructure/          # DB implementation (Task 2)
│   └── web/                     # HTTP endpoints (Tasks 5-7)
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Prerequisites

- Go 1.22+

## Build & Run

```bash
# Build the binary
make build

# Run the server (default port 8080)
make run

# Run with custom port
PORT=3000 make run

# Run tests
make test

# Run linter
make lint
```

## Health Check

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

## Migration Roadmap

1. **Task 1** (this PR): Go project skeleton, domain types, and DTOs
2. **Task 2**: Database layer (repository interfaces and implementations)
3. **Task 3**: Application layer (use cases / services)
4. **Task 4**: Infrastructure (EF Core equivalent with GORM or sqlx)
5. **Tasks 5-7**: HTTP endpoints and middleware
