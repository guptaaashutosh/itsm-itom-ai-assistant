# ITSM-ITOM AI Assistant

An AI-powered incident triage assistant for ITSM/ITOM SaaS platforms, built in Go. It uses Gin for HTTP, pgx for PostgreSQL, Viper for config, and dbmate for migrations. The system exposes a POST /incidents endpoint that uses basic AI logic to assign severity and category, stores the incident, and returns the result.

## Features
- Modular, idiomatic Go project structure
- Gin-based REST API
- PostgreSQL with pgx
- Viper for config management
- dbmate for migrations
- Simple AI logic for incident triage

## Project Structure

```
itsm-itom-ai-assistant/
│
├── cmd/                    # Application entry points (main.go)
│
├── config/                 # Configuration files and logic
│
├── constants/              # Project-wide constant values
│
├── dal/                    # Data Access Layer (database queries)
│
├── db/
│   └── migrations/         # Database migration scripts
│
├── internal/               # Internal utilities and helper functions
│
├── middlewares/            # Gin middlewares (e.g., error recovery, request logging)
│
├── repository/             # Repository pattern implementations for data persistence
│
├── resource/               # External resource connections (e.g., DB, cache)
│
├── restapiv1/
│   ├── controllers/        # HTTP handler functions (controllers)
│   ├── models/
│   │   ├── dto/            # Data Transfer Objects
│   │   ├── request/        # Request payload definitions
│   │   └── response/       # Response payload definitions
│   ├── routes/             # API route definitions
│   └── validations/        # Request validation logic
│
├── swagger/                # OpenAPI/Swagger documentation
│
├── Dockerfile              # Docker container configuration
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
└── README.md               # Project documentation
```

- `.config/`: Environment variables (`.env`).
- `cmd/`: Application entrypoint (`main.go`).
- `config/`: Loads configuration using Viper.
- `dal/`: Data access layer (raw SQL/query builder).
- `db/`: Database migrations.
- `internal/`: Internal utilities/helpers.
- `middlewares/`: HTTP middlewares.
- `repository/`: Repository implementations.
- `resource/`: Connection setup (e.g., DB).
- `restapiv1/`: API layer (constants, controllers, models, routes, validations).
- `swagger/`: OpenAPI specs.
- `Dockerfile`: Build/run the Go app.
- `.env.example`: Example environment variables.

## Setup
1. Copy `.env.example` to `.config/.env` and update values.
2. Run `go mod tidy` to install dependencies.
3. Run migrations: `dbmate --url="postgres://DB_USER:DB_PASSWORD@DB_HOST:DB_PORT/DB_NAME?sslmode=disable" up`
4. Start the server: `go run ./cmd/main.go`

## API
### POST /incidents
**Request:**
```
{
  "title": "string",
  "description": "string",
  "affected_service": "string"
}
```
**Response:**
```
{
  "id": "uuid",
  "title": "string",
  "description": "string",
  "affected_service": "string",
  "ai_severity": "string",
  "ai_category": "string",
  "created_at": "timestamp"
}
```

## AI Logic
- If the title or description contains "outage": severity = critical, category = infrastructure
- If contains "fail" or "error": severity = high, category = application
- Otherwise: severity = normal, category = general

## Cursor Usage
This project was scaffolded and iteratively built using [Cursor](https://www.cursor.so/), an AI-powered IDE for modern development workflows.

## Usage

- All business logic must be outside `main.go`.
- Configuration is loaded via `config/config.go` using Viper.
- See each directory for more details. 