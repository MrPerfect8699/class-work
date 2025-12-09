# Hexagonal Architecture Refactoring Summary

## Overview

Your Classwork API has been successfully refactored into a **Hexagonal Architecture** (also known as Ports & Adapters). This architecture style provides better separation of concerns, improved testability, and enhanced maintainability.

## Architecture Layers

```
┌─────────────────────────────────────────────────────────────┐
│                    EXTERNAL ADAPTERS                        │
├──────────────────────────┬──────────────────────────────────┤
│   HTTP (Chi Framework)   │    MongoDB Persistence          │
│  ┌────────────────────┐  │  ┌──────────────────────────┐   │
│  │ Handlers           │  │  │ Repositories             │   │
│  │ Server             │  │  │ Connection Management    │   │
│  │ Middleware         │  │  │ Query Building           │   │
│  └────────────────────┘  │  └──────────────────────────┘   │
├──────────────────────────┴──────────────────────────────────┤
│                    PORTS (Interfaces)                       │
│  ┌──────────────────────┐      ┌──────────────────────┐    │
│  │ AuthService          │      │ HomeworkService      │    │
│  └──────────────────────┘      └──────────────────────┘    │
├──────────────────────────────────────────────────────────────┤
│               APPLICATION SERVICES LAYER                    │
│  ┌──────────────────────┐      ┌──────────────────────┐    │
│  │ AuthServiceImpl       │      │ HomeworkServiceImpl   │    │
│  │ (Business Logic)     │      │ (Business Logic)     │    │
│  └──────────────────────┘      └──────────────────────┘    │
├──────────────────────────────────────────────────────────────┤
│                   DOMAIN LAYER (Core)                       │
│  ┌────────────────────┐         ┌────────────────────┐     │
│  │ Teacher            │         │ Homework           │     │
│  │ TeacherRepository  │         │ HomeworkRepository │     │
│  │ (Interfaces)       │         │ (Interfaces)       │     │
│  └────────────────────┘         └────────────────────┘     │
├──────────────────────────────────────────────────────────────┤
│              CONFIGURATION LAYER                            │
│              (Environment & Settings)                       │
└──────────────────────────────────────────────────────────────┘
```

## Directory Structure

```
backend/
├── config/                          # Configuration management
│   └── config.go                   # Config loader for env vars
│
├── internal/                        # Internal packages (not exported)
│   ├── domain/                     # Core business entities & interfaces
│   │   ├── teacher.go              # Teacher entity
│   │   └── homework.go             # Homework entity
│   │
│   ├── application/                # Application services (use cases)
│   │   ├── auth_service.go         # Authentication logic
│   │   └── homework_service.go     # Homework management logic
│   │
│   ├── ports/                      # Port interfaces (contracts)
│   │   ├── auth.go                 # AuthService interface
│   │   └── homework.go             # HomeworkService interface
│   │
│   └── adapters/                   # Concrete implementations
│       ├── http/                   # HTTP adapter
│       │   ├── server.go           # HTTP server setup
│       │   ├── handler.go          # Route registration
│       │   ├── auth_handler.go     # Auth request handlers
│       │   ├── homework_handler.go # Homework request handlers
│       │   └── middleware.go       # Auth middleware
│       │
│       └── persistence/            # Database adapter
│           ├── mongodb.go          # MongoDB connection
│           ├── teacher_repository.go
│           └── homework_repository.go
│
├── pkg/openapi/                    # OpenAPI specification
│   └── openapi.yaml                # Full API documentation
│
├── main.go                         # Application entry point
│
├── .env                            # Environment configuration
├── .env.example                    # Configuration template
├── README.md                       # Comprehensive documentation
└── go.mod                          # Go module definition
```

## Key Benefits of Hexagonal Architecture

### 1. **Separation of Concerns**
- Business logic (domain & application) is completely isolated
- Infrastructure concerns (HTTP, database) are pluggable
- Easy to understand and modify each layer independently

### 2. **Testability**
- Domain and application layers have no external dependencies
- Easy to mock repositories and services for unit tests
- Can test business logic without database or HTTP

### 3. **Flexibility**
- Swap HTTP adapter (Gin → Echo, Fiber, etc.) without touching business logic
- Replace MongoDB with PostgreSQL by implementing the same repository interfaces
- Reuse services across different interfaces (CLI, gRPC, etc.)

### 4. **Maintainability**
- Clear dependencies flow (inward)
- Easy to locate code for specific features
- Reduced coupling between components

### 5. **Scalability**
- Can evolve different layers independently
- Easy to add new adapters for new requirements
- Clean interfaces make it simple to add features

## Layer Responsibilities

### Domain Layer (`internal/domain/`)
- **Entities**: Pure business objects (Teacher, Homework)
- **Interfaces**: Repository contracts
- **No dependencies**: Never imports from other layers
- **Purpose**: Represent business rules and constraints

### Application Layer (`internal/application/`)
- **Services**: Implement business use cases
- **No HTTP/Database logic**: Abstracted through interfaces
- **Depends on**: Domain and Ports layers
- **Purpose**: Orchestrate domain logic to fulfill user requirements

### Ports Layer (`internal/ports/`)
- **Interfaces**: Define contracts for services
- **No implementation**: Just contracts
- **Purpose**: Define what the application exposes to external world

### Adapters Layer (`internal/adapters/`)
- **HTTP Adapter**: Handles incoming requests, translates to domain objects
- **Persistence Adapter**: Implements repository interfaces, handles database operations
- **No business logic**: Just translation and integration
- **Purpose**: Connect application to external systems

### Configuration Layer (`config/`)
- **Settings**: Load and manage environment variables
- **Logging**: Log configuration for debugging
- **Purpose**: Centralized configuration management

## Configuration Management

The application uses environment variables with `config.go`:

```go
// Automatically loaded from .env file
config, _ := config.LoadConfig()

// Access configuration
cfg.Server.Port        // 8080
cfg.Database.URI       // MongoDB connection string
cfg.JWT.Secret         // JWT signing secret
cfg.JWT.ExpirationTime // Token expiration time
```

All configuration can be overridden via environment variables:

```bash
# Server
ENV=production
HOST=0.0.0.0
PORT=8080
READ_TIMEOUT=10s
WRITE_TIMEOUT=10s
IDLE_TIMEOUT=120s

# Database
MONGODB_URI=mongodb+srv://...
DATABASE_NAME=classwork
DB_CONN_TIMEOUT=10s
DB_MAX_POOL_SIZE=100
DB_MIN_POOL_SIZE=10

# JWT
JWT_SECRET=your-secret
JWT_EXPIRATION=72h
JWT_REFRESH_TIME=24h
```

## OpenAPI Documentation

The API is fully documented in `pkg/openapi/openapi.yaml`:

- Complete endpoint specifications
- Request/response schemas
- Authentication details
- Error responses
- Security requirements

View the documentation at: https://editor.swagger.io/ (paste YAML content)

## HTTP Server Management

The `internal/adapters/http/server.go` provides:

- Clean server lifecycle management
- Graceful shutdown with timeout
- Configurable timeouts
- Error handling
- Chi router with integrated middleware
- CORS support for localhost:*

```go
// Create server
server := http.NewServer(handler, &cfg.Server)

// Start (blocks until shutdown signal)
if err := server.Start(); err != nil {
    log.Fatal(err)
}

// Graceful shutdown
server.Shutdown(ctx)
```

## HTTP Framework Details

The application uses **Chi** router with the following features:

### Middleware Stack
- `middleware.Logger` - Request logging
- `middleware.Recoverer` - Panic recovery
- `middleware.RequestID` - Request ID tracking
- `cors.Handler` - CORS support for localhost:*

### Authentication Middleware
- JWT validation via `Authorization: Bearer <token>` header
- Automatic teacher ID extraction from token
- Context-based passing of authenticated user data

### Handler Pattern
All handlers follow the standard Go HTTP pattern:
```go
func (h *Handler) CreateHomework(w http.ResponseWriter, r *http.Request) {
    // Read JSON from r.Body using json.NewDecoder
    // Extract URL parameters using chi.URLParam(r, "id")
    // Get context values from r.Context().Value("teacherID")
    // Write responses using w.Header().Set() and w.WriteHeader()
}
```

## Running the Application

```bash
# Development
go run .

# Production build
go build -o classwork
./classwork

# With environment override
PORT=9000 go run .
```

## Testing Strategy

### Unit Tests
- Test services with mocked repositories
- Test domain logic without database
- Fast and isolated

### Integration Tests
- Test repositories with real MongoDB
- Test HTTP handlers with real services
- Slower but comprehensive

### Example Test Implementation

```go
// Mock repository for testing
type MockTeacherRepo struct {
    SaveFunc func(*domain.Teacher) error
}

func (m *MockTeacherRepo) Save(t *domain.Teacher) error {
    return m.SaveFunc(t)
}

// Test service with mock
func TestAuthService(t *testing.T) {
    mockRepo := &MockTeacherRepo{...}
    authSvc := application.NewAuthService(mockRepo, cfg)
    // Test logic
}

// Test HTTP handler
func TestCreateHomeworkHandler(t *testing.T) {
    router := chi.NewRouter()
    handler := http.NewHandler(authSvc, homeworkSvc)
    handler.RegisterRoutes(router)
    
    // Make test request
    req := httptest.NewRequest("POST", "/api/homework", body)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    // Assert response
}
```

## Adding New Features

### Example: Add a new "Student" domain

1. **Create domain entity** (`internal/domain/student.go`)
   ```go
   type Student struct {
       ID    primitive.ObjectID
       Name  string
       Email string
   }
   
   type StudentRepository interface {
       Save(*Student) error
       FindByID(interface{}) (*Student, error)
   }
   ```

2. **Create port interface** (`internal/ports/student.go`)
   ```go
   type StudentService interface {
       RegisterStudent(*Student) error
       GetStudent(id interface{}) (*Student, error)
   }
   ```

3. **Create service** (`internal/application/student_service.go`)
   ```go
   type StudentServiceImpl struct {
       studentRepo domain.StudentRepository
   }
   ```

4. **Create repository** (`internal/adapters/persistence/student_repository.go`)
   ```go
   type StudentRepositoryImpl struct {
       db *mongo.Database
   }
   ```

5. **Create handlers** (`internal/adapters/http/student_handler.go`)
   ```go
   func (h *Handler) RegisterStudent(w http.ResponseWriter, r *http.Request) { ... }
   ```

6. **Register routes** in `internal/adapters/http/handler.go`

## Framework Migration: Gin to Chi

The project has been successfully migrated from Gin to Chi framework:
- **Chi Router** provides better modularity and CORS support
- **Standard HTTP patterns** using `http.ResponseWriter` and `*http.Request`
- **Chi middleware** for request logging, recovery, and request ID tracking
- **CORS middleware** from `go-chi/cors` for localhost:* support

### Files Removed During Migration
- `db.go` - Replaced by `internal/adapters/persistence/mongodb.go`
- `handlers.go` - Replaced by `internal/adapters/http/*_handler.go`
- `middleware.go` - Replaced by `internal/adapters/http/middleware.go`
- `models.go` - Replaced by `internal/domain/*.go`
- `config.go` (old) - Replaced by `config/config.go`
- `chi_routes.go` - Deprecated (routes registered via `handler.RegisterRoutes()`)

## Next Steps

1. **Add Unit Tests**: Create `*_test.go` files for services
2. **Add Integration Tests**: Test repositories with real database
3. **Add Error Handling**: Create custom error types in domain layer
4. **Add Logging**: Use structured logging (zap, logrus)
5. **Add Validation**: Add input validation in service layer
6. **Add More Endpoints**: Follow the pattern for new features
7. **Deploy**: Build binary and run in production

## Summary

Your project is now structured using **hexagonal architecture**, providing:
✅ Clear separation of concerns
✅ Improved testability
✅ Better maintainability
✅ Enhanced flexibility
✅ Production-ready structure
✅ Full API documentation
✅ Centralized configuration
✅ Graceful server lifecycle management

The application is running successfully on port 8080 and ready for testing and further development!
