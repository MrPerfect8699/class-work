# Classwork API - Hexagonal Architecture

This is a RESTful API for managing classroom homework assignments, built with Go using hexagonal (ports and adapters) architecture.

## Architecture

The project follows **Hexagonal Architecture (Ports & Adapters)** pattern:

```
├── cmd/                          # Application entry points
│   └── main.go                   # Application entry point
├── config/                       # Configuration management
│   └── config.go                 # Configuration loader
├── internal/                     # Internal packages (not exported)
│   ├── domain/                   # Core business logic entities
│   │   ├── teacher.go           # Teacher entity and repository interface
│   │   └── homework.go          # Homework entity and repository interface
│   ├── application/             # Application services
│   │   ├── auth_service.go      # Authentication service
│   │   └── homework_service.go  # Homework management service
│   ├── ports/                   # Port interfaces (contracts)
│   │   ├── auth.go              # Authentication service interface
│   │   └── homework.go          # Homework service interface
│   └── adapters/                # Concrete implementations (adapters)
│       ├── http/                # HTTP adapter (Gin framework)
│       │   ├── server.go        # HTTP server setup
│       │   ├── handler.go       # Route handler setup
│       │   ├── auth_handler.go  # Authentication handlers
│       │   ├── homework_handler.go # Homework handlers
│       │   └── middleware.go    # HTTP middleware
│       └── persistence/         # Persistence adapter (MongoDB)
│           ├── mongodb.go       # MongoDB connection
│           ├── teacher_repository.go # Teacher persistence
│           └── homework_repository.go # Homework persistence
└── pkg/openapi/                 # OpenAPI specification
    └── openapi.yaml             # OpenAPI 3.0 specification
```

## Key Concepts

### Domain Layer
- **Entities**: `Teacher`, `Homework`, `Submission`
- **Interfaces**: `TeacherRepository`, `HomeworkRepository`
- Pure business logic without external dependencies

### Application Layer
- **Services**: `AuthServiceImpl`, `HomeworkServiceImpl`
- Implements business use cases
- Uses domain entities and repositories

### Ports Layer
- **Interfaces**: `AuthService`, `HomeworkService`
- Defines contracts for the application layer

### Adapters Layer
- **HTTP Adapter**: Gin-based REST API implementation
- **Persistence Adapter**: MongoDB implementation
- Concrete implementations that interact with external systems

## Features

- ✅ Teacher authentication with JWT
- ✅ Create, read, update, delete homework assignments
- ✅ Role-based access control (protected routes)
- ✅ Password hashing with bcrypt
- ✅ MongoDB persistence
- ✅ Configuration management
- ✅ Comprehensive error handling
- ✅ OpenAPI 3.0 specification

## Prerequisites

- Go 1.19+
- MongoDB 4.0+
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourname/classwork.git
cd classwork/backend
```

2. Install dependencies:
```bash
go mod download
go mod tidy
```

3. Create a `.env` file based on `.env.example`:
```bash
cp .env.example .env
```

4. Update `.env` with your configuration (especially MongoDB URI and JWT secret)

## Running the Application

### Development Mode

```bash
go run .
```

The server will start on `http://localhost:8080` by default.

### Building for Production

```bash
go build -o classwork
./classwork
```

## API Endpoints

### Authentication
- `POST /api/register` - Register a new teacher
- `POST /api/login` - Login and get JWT token

### Homework Management (Protected)
- `POST /api/homework` - Create a new homework assignment
- `GET /api/homeworks` - List all homework assignments
- `GET /api/homework/:id` - Get a specific homework assignment
- `PUT /api/homework/:id` - Update a homework assignment
- `DELETE /api/homework/:id` - Delete a homework assignment

### Health Check
- `GET /health` - Check server health status

## Configuration

Configuration is managed through environment variables. See `.env.example` for all available options:

- `ENV` - Application environment (development, staging, production)
- `PORT` - Server port (default: 8080)
- `MONGODB_URI` - MongoDB connection string
- `JWT_SECRET` - Secret key for JWT signing
- `JWT_EXPIRATION` - JWT token expiration time (default: 72h)

## API Documentation

OpenAPI/Swagger documentation is available in `pkg/openapi/openapi.yaml`.

To view the documentation:
1. Copy the content of `openapi.yaml`
2. Visit [swagger.io/tools/swagger-editor](https://editor.swagger.io/)
3. Paste the YAML content

## Development

### Code Structure
- **Domain-driven**: Business logic is completely decoupled from infrastructure
- **Testable**: All services have clear interfaces for mocking
- **Maintainable**: Clear separation of concerns

### Adding New Features

1. Create domain entity in `internal/domain/`
2. Create repository interface in the entity file
3. Create service in `internal/application/`
4. Create repository implementation in `internal/adapters/persistence/`
5. Create HTTP handlers in `internal/adapters/http/`
6. Register routes in `internal/adapters/http/handler.go`

## Error Handling

The application uses structured error handling:
- Domain layer returns domain-specific errors
- Application layer wraps errors with context
- HTTP adapter translates errors to appropriate HTTP status codes

## Security

- Passwords are hashed using bcrypt
- JWT tokens include expiration and claims
- Protected routes require valid bearer token
- All inputs are validated

## Testing

Run tests with:
```bash
go test ./...
```

For coverage:
```bash
go test -cover ./...
```

## Troubleshooting

### MongoDB Connection Issues
- Check `MONGODB_URI` is correct
- Ensure MongoDB is running and accessible
- Verify network connectivity and firewall rules

### JWT Token Errors
- Ensure `JWT_SECRET` is set in `.env`
- Check token hasn't expired
- Verify token format: `Bearer <token>`

### CORS Issues
- Configure CORS middleware in `internal/adapters/http/middleware.go` if needed

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see LICENSE file for details.

## Support

For issues and questions, please create an issue on GitHub.
