# Development Guide

## Quick Start

### Prerequisites
- Go 1.19+
- MongoDB 4.0+ or MongoDB Atlas account
- Git

### Setup

```bash
# Clone and navigate to backend
cd classwork/backend

# Install dependencies
go mod download

# Create .env from template
cp .env.example .env

# Update .env with your settings (especially MONGODB_URI and JWT_SECRET)
```

### Running Locally

```bash
# Development mode
go run .

# With hot reload (install air first: go install github.com/cosmtrek/air@latest)
air

# Production build
go build -o classwork
./classwork
```

The server will be available at `http://localhost:8080`

## Project Structure Overview

### Key Directories

| Directory | Purpose |
|-----------|---------|
| `config/` | Configuration loading and management |
| `internal/domain/` | Core business entities and interfaces |
| `internal/application/` | Business logic implementation |
| `internal/ports/` | Service interfaces/contracts |
| `internal/adapters/http/` | REST API implementation |
| `internal/adapters/persistence/` | Database implementation |
| `pkg/openapi/` | API documentation |

### Layer Dependencies

```
HTTP Adapter → Application Services → Domain Entities
   ↓              ↓                        ↓
Config ← Persistence Adapter ← Repository Interfaces
```

**Rule**: Dependencies flow inward. Domain layer depends on nothing.

## Common Tasks

### Testing an Endpoint

#### 1. Register a Teacher
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepass123"
  }'
```

Response:
```json
{
  "id": "507f1f77bcf86cd799439011"
}
```

#### 2. Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepass123"
  }'
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### 3. Create Homework (Protected - requires token)
```bash
curl -X POST http://localhost:8080/api/homework \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "title": "Math Assignment",
    "description": "Complete exercises 1-10",
    "className": "Class 10-A",
    "subject": "Mathematics",
    "attachments": "https://example.com/assignment.pdf"
  }'
```

#### 4. Get Homeworks
```bash
curl -X GET http://localhost:8080/api/homeworks \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

#### 5. Health Check
```bash
curl http://localhost:8080/health
```

### Adding a New Endpoint

#### Step 1: Extend Domain (if needed)
Add entity in `internal/domain/`:
```go
// internal/domain/submission.go
package domain

type Submission struct {
    ID        primitive.ObjectID
    HomeworkID primitive.ObjectID
    StudentName string
}

type SubmissionRepository interface {
    Save(*Submission) error
    FindByHomeworkID(id interface{}) ([]Submission, error)
}
```

#### Step 2: Create Service Interface
```go
// internal/ports/submission.go
package ports

type SubmissionService interface {
    CreateSubmission(*domain.Submission) error
    GetSubmissions(homeworkID interface{}) ([]domain.Submission, error)
}
```

#### Step 3: Implement Service
```go
// internal/application/submission_service.go
package application

type SubmissionServiceImpl struct {
    submissionRepo domain.SubmissionRepository
}

func NewSubmissionService(repo domain.SubmissionRepository) *SubmissionServiceImpl {
    return &SubmissionServiceImpl{submissionRepo: repo}
}

func (s *SubmissionServiceImpl) CreateSubmission(sub *domain.Submission) error {
    // Business logic here
    return s.submissionRepo.Save(sub)
}
```

#### Step 4: Implement Repository
```go
// internal/adapters/persistence/submission_repository.go
package persistence

type SubmissionRepositoryImpl struct {
    db *mongo.Database
}

func (r *SubmissionRepositoryImpl) Save(sub *domain.Submission) error {
    // MongoDB operations
}
```

#### Step 5: Create Handlers
```go
// internal/adapters/http/submission_handler.go
package http

func (h *Handler) CreateSubmission(c *gin.Context) {
    var req CreateSubmissionRequest
    // Handle request
}

func (h *Handler) GetSubmissions(c *gin.Context) {
    // Handle request
}
```

#### Step 6: Register Routes
```go
// internal/adapters/http/handler.go
func (h *Handler) RegisterRoutes(r *gin.Engine) {
    // ... existing routes ...
    
    protected := r.Group("/api")
    protected.Use(h.AuthMiddleware())
    {
        protected.POST("/submission", h.CreateSubmission)
        protected.GET("/submissions/:homeworkId", h.GetSubmissions)
    }
}
```

#### Step 7: Update Main
```go
// main.go
submissionRepo := persistence.NewSubmissionRepository(mongoDB.GetDB())
submissionService := application.NewSubmissionService(submissionRepo)
handler := http.NewHandler(authService, homeworkService, submissionService)
```

### Writing Tests

#### Service Test Example
```go
// internal/application/submission_service_test.go
package application

import "testing"

func TestCreateSubmission(t *testing.T) {
    // Mock repository
    mockRepo := &MockSubmissionRepository{
        SaveFunc: func(s *domain.Submission) error {
            return nil
        },
    }
    
    service := NewSubmissionService(mockRepo)
    
    // Test
    err := service.CreateSubmission(&domain.Submission{...})
    
    if err != nil {
        t.Errorf("expected no error, got %v", err)
    }
}
```

#### Running Tests
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -run TestCreateSubmission ./internal/application/
```

### Debugging

#### Using Delve Debugger
```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Start debugging
dlv debug

# In delve prompt:
# (dlv) break main.main
# (dlv) continue
# (dlv) next
# (dlv) print variable_name
# (dlv) exit
```

#### Using VS Code Debugger
Add to `.vscode/launch.json`:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}",
            "env": {},
            "args": [],
            "showLog": true
        }
    ]
}
```

### Environment Variables

Key variables:
- `ENV` - Application mode (development, staging, production)
- `PORT` - Server port
- `MONGODB_URI` - MongoDB connection string
- `JWT_SECRET` - JWT signing key
- `READ_TIMEOUT`, `WRITE_TIMEOUT` - Request timeouts

Override at runtime:
```bash
JWT_SECRET=new-secret PORT=9000 go run .
```

### Database Management

#### MongoDB Commands
```bash
# Connect to local MongoDB
mongosh

# List databases
show databases

# Use classwork database
use classwork

# Show collections
show collections

# Query teachers
db.teachers.find()

# Query homework
db.homework.find()

# Count documents
db.teachers.countDocuments()

# Drop collection
db.teachers.drop()
```

### Performance Tips

1. **Connection Pooling**: MongoDB driver automatically manages connection pool
   - `DB_MAX_POOL_SIZE=100` (default)
   - `DB_MIN_POOL_SIZE=10` (default)

2. **Timeout Settings**:
   - Adjust `READ_TIMEOUT`, `WRITE_TIMEOUT` based on your needs
   - Default: 10 seconds (appropriate for most cases)

3. **Database Indexes**:
   Add indexes for frequently queried fields:
   ```go
   // In persistence layer initialization
   indexModel := mongo.IndexModel{
       Keys: bson.D{
           {Key: "email", Value: 1},
       },
       Options: options.Index().SetUnique(true),
   }
   db.Collection("teachers").Indexes().CreateOne(context.TODO(), indexModel)
   ```

### Code Style

Follow Go conventions:
- `gofmt` - Format code
- `golint` - Lint code
- `go vet` - Vet code

```bash
# Format all files
go fmt ./...

# Run linter (install golangci-lint first)
golangci-lint run
```

### Common Issues

#### "MongoDB connection refused"
- Ensure MongoDB is running
- Check `MONGODB_URI` is correct
- Verify network/firewall settings

#### "Invalid token"
- Check token hasn't expired
- Verify `JWT_SECRET` matches between login and auth
- Ensure Bearer token format: `Bearer <token>`

#### "Email already taken"
- Clear `teachers` collection or use different email
- Check for duplicates: `db.teachers.find({email: "test@example.com"})`

#### "CORS Error"
- Add CORS middleware in `internal/adapters/http/middleware.go`
- Or configure frontend to include proper headers

## Useful Resources

- [Go Documentation](https://golang.org/doc)
- [MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [JWT-Go](https://github.com/dgrijalva/jwt-go)
- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)

## Getting Help

1. Check error messages carefully - they're usually descriptive
2. Use `go vet ./...` to find potential issues
3. Check logs with `go run . 2>&1 | grep -i error`
4. Enable verbose logging: `export GODEBUG=... `
5. Use debugger for complex issues

## Next Development Areas

- [ ] Add email validation
- [ ] Add pagination to list endpoints
- [ ] Add submission tracking
- [ ] Add grade management
- [ ] Add file upload support
- [ ] Add request logging middleware
- [ ] Add metrics/monitoring
- [ ] Add database transactions for complex operations
- [ ] Add caching layer
- [ ] Add API versioning

Good luck with your development! 🚀
