# go-api

A Go port of the Clean Architecture .NET solution in this repo, using Go's
standard project layout conventions.

## Layout

```
go-api/
├── cmd/api/main.go           # entrypoint (equivalent of src/Web/Program.cs)
├── internal/
│   ├── domain/               # equivalent of src/Domain/
│   │   ├── entities/
│   │   ├── events/
│   │   ├── enums/
│   │   ├── valueobjects/
│   │   └── errors/
│   ├── application/          # equivalent of src/Application/
│   │   ├── todoitems/
│   │   ├── todolists/
│   │   ├── weatherforecasts/
│   │   ├── common/
│   │   │   ├── interfaces/
│   │   │   ├── models/
│   │   │   └── validation/
│   │   └── middleware/       # replaces MediatR behaviours
│   ├── infrastructure/       # equivalent of src/Infrastructure/
│   │   ├── persistence/
│   │   ├── identity/
│   │   └── config/
│   └── web/                  # equivalent of src/Web/
│       ├── endpoints/
│       ├── middleware/
│       └── problemdetails/
├── go.mod
├── go.sum
└── Makefile
```

## Mapping from .NET to Go

| .NET (src/)                              | Go (go-api/)                                      |
| ---------------------------------------- | ------------------------------------------------- |
| `src/Web/Program.cs`                     | `cmd/api/main.go`                                 |
| `src/Domain/**`                          | `internal/domain/**`                              |
| `src/Application/**`                     | `internal/application/**`                         |
| MediatR pipeline behaviours              | `internal/application/middleware/`                |
| FluentValidation                         | `github.com/go-playground/validator/v10`          |
| EF Core                                  | `github.com/jmoiron/sqlx`                         |
| ASP.NET Identity (bearer tokens)         | `github.com/golang-jwt/jwt/v5` + `golang.org/x/crypto/bcrypt` |
| `src/Web/Endpoints/**`                   | `internal/web/endpoints/`                         |
| ASP.NET middleware                       | `internal/web/middleware/`                        |
| ProblemDetails                           | `internal/web/problemdetails/`                    |

## Quickstart

```sh
make build
make run
curl http://localhost:8080/health
```
