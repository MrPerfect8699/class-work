# 📚 Classwork API - Hexagonal Architecture Edition 🏗️

> _Where homework gets organized, teachers get caffeinated, ☕ and code stays clean!_

This is a **RESTful API** for managing classroom homework assignments, built with **Go** using hexagonal (ports and adapters) architecture. It's like a Swiss Army knife, but for teachers who actually understand clean code! 🎯

---

## 🎬 The Setup (Starring Our Beautiful Architecture)

The project follows **Hexagonal Architecture (Ports & Adapters)** - think of it as separating the business folks from the IT folks, but they're all friends! 🤝

```
📦 Backend Project Structure 
├── 🎯 config/                          # Configuration (the rulebook)
│   └── ⚙️  config.go                  
├── 🏛️  internal/                      # The Secret Sauce™
│   ├── 🎭 domain/                     # Core business logic (the stars)
│   │   ├── 👨‍🏫 teacher.go              # Teacher entity and repository interface
│   │   └── 📝 homework.go             # Homework entity and repository interface
│   │
│   ├── 🚀 application/                # Application services (the stage)
│   │   ├── 🔐 auth_service.go         # "Who are you?" service
│   │   └── 📚 homework_service.go     # "Let's manage homework!" service
│   │
│   ├── 🚪 ports/                      # Port interfaces (the contract)
│   │   ├── 🔑 auth.go                 # What authentication should do
│   │   └── 📋 homework.go             # What homework management should do
│   │
│   └── 🔌 adapters/                   # Concrete implementations (the muscle)
│       ├── 🌐 http/                   # HTTP adapter (our web gateway)
│       │   ├── 🖥️  server.go          # HTTP server setup
│       │   ├── 🛣️  handler.go         # Route handler setup
│       │   ├── 🔓 auth_handler.go     # Login/Register endpoints
│       │   ├── 📤 homework_handler.go # CRUD for homework
│       │   ├── 🛡️  middleware.go      # Security checks
│       │   └── 🌍 cors.go             # CORS config for localhost
│       │
│       └── 💾 persistence/            # Database adapter (permanent storage)
│           ├── 🗄️  mongodb.go         # MongoDB connection handler
│           ├── 👨‍💼 teacher_repository.go    # Save/fetch teachers
│           └── 📖 homework_repository.go   # Save/fetch homework
│
└── 📖 pkg/openapi/                    # API Documentation (the manual)
    └── 📄 openapi.yaml                # Full OpenAPI 3.0 spec
```

---

## 🎭 The Cast of Characters

### 👥 Domain Layer - _"The Stars of the Show"_
- **🎬 Entities**: `Teacher` (the boss), `Homework` (the task)
- **📋 Interfaces**: `TeacherRepository`, `HomeworkRepository` (the scripts)
- **✨ Magic**: Pure business logic, zero external dependencies = **zero drama!**

### 🚪 Ports Layer - _"The Bouncer at the Club"_
- **🔐 Interfaces**: `AuthService`, `HomeworkService`
- **📝 Job**: Defines what the application PROMISES to do (the contract)
- **💼 Style**: Professional, clean, business-like

### 🔌 Adapters Layer - _"The Stunt Doubles"_
- **🌐 HTTP Adapter**: Chi Router + Middleware = _"Hello Internet!"_
- **💾 Persistence Adapter**: MongoDB = _"Remember everything, forget nothing!"_
- **🤝 Purpose**: Translates between the outside world and our beautiful domain

---

## ✨ Features (The Cool Stuff)

| Feature | Status | Emoji |
|---------|--------|-------|
| Teacher authentication with JWT | ✅ Implemented | 🔐 |
| Homework CRUD operations | ✅ Implemented | 📝 |
| Role-based access control | ✅ Implemented | 🛡️  |
| Password hashing (bcrypt) | ✅ Implemented | 🔒 |
| MongoDB persistence | ✅ Implemented | 💾 |
| Configuration management | ✅ Implemented | ⚙️  |
| Error handling (awesome) | ✅ Implemented | 🎯 |
| OpenAPI 3.0 documentation | ✅ Implemented | 📖 |
| CORS for localhost:* | ✅ Implemented | 🌍 |
| Chi Router (modern & clean) | ✅ Implemented | ⚡ |

---

## 🛠️ Prerequisites (What You Need in Your Toolkit)

- **Go** 1.19+ _(the brains)_
- **MongoDB** 4.0+ _(the memory)_
- **Git** _(the time machine)_
- **Coffee** ☕ _(optional but highly recommended)_

---

## 📦 Installation (Getting Started)

### Step 1️⃣ - Clone the Repository
```bash
git clone https://github.com/yourname/classwork.git
cd classwork/backend
```

### Step 2️⃣ - Install Dependencies
```bash
go mod download
go mod tidy
```

### Step 3️⃣ - Setup Configuration
```bash
cp .env.example .env
```

### Step 4️⃣ - Configure Environment
Edit `.env` with your settings:
```env
# 🔓 Server
HOST=localhost
PORT=8080

# 💾 MongoDB
MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/

# 🔑 JWT
JWT_SECRET=your-super-secret-key-here
JWT_EXPIRATION=72h
```

---

## 🚀 Running the Application (Let's Go!)

### Development Mode 🎮
```bash
go run .
```

### Production Build 📦
```bash
go build -o classwork .
./classwork
```

### With Custom Port 🎯
```bash
PORT=9000 go run .
```

The API will start on `http://localhost:8080` 🎉

---

## 📡 API Endpoints (The Menu)

### 🔓 Public Endpoints (No Password Needed)

```bash
# Register a new teacher
POST /api/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "secret123"
}

# Login and get JWT token
POST /api/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "secret123"
}
```

### 🔐 Protected Endpoints (Bring Your Token!)

```bash
# Headers required
Authorization: Bearer <your-jwt-token-here>

# Create homework assignment
POST /api/homework
{
  "title": "Math Assignment",
  "description": "Chapter 5 problems",
  "className": "Period 1",
  "subject": "Mathematics"
}

# Get all your homework
GET /api/homeworks

# Get specific homework
GET /api/homework/{id}

# Update homework
PUT /api/homework/{id}
{
  "title": "Updated Title"
}

# Delete homework
DELETE /api/homework/{id}
```

### 💚 Health Check (Am I Alive?)
```bash
GET /health
```

---

## 📚 Project Structure Deep Dive

### 🎯 Domain Layer
_The Heart of the Business_
- **What lives here**: Core entities and repository interfaces
- **Who can touch it**: ONLY application services
- **What it knows**: Nothing about HTTP, databases, or the outside world
- **Personality**: Pure, innocent, untouched by framework drama

### 🚀 Application Layer
_The Brain of the Operation_
- **What lives here**: Business logic and use case implementations
- **Who can touch it**: HTTP adapters and persistence adapters
- **What it knows**: Domains and ports, but NOT the outside world
- **Personality**: Smart, decisive, knows what needs to happen

### 🌐 HTTP Adapter
_The Spokesperson to the Internet_
- **Framework**: Chi Router (lightweight, clean, Go-native) ⚡
- **Job**: Convert HTTP requests → domain objects → HTTP responses
- **Middleware**: JWT validation, CORS, request logging
- **Personality**: Chatty, friendly, handles all the web stuff

### 💾 Persistence Adapter
_The Librarian of the Application_
- **Database**: MongoDB (flexible, scalable, document-based)
- **Job**: Save and retrieve data from the database
- **Implementation**: Implements repository interfaces
- **Personality**: Quiet, reliable, never forgets

---

## 🧪 Testing the API (Let's Try It Out!)

### Using cURL 🌐

```bash
# 1️⃣ Register
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "email": "jane@school.com",
    "password": "password123"
  }'

# 2️⃣ Login (save the token!)
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "jane@school.com",
    "password": "password123"
  }'
# Response: {"token":"eyJhbGciOiJIUzI1NiIs..."}

# 3️⃣ Create homework (replace TOKEN with your actual token)
curl -X POST http://localhost:8080/api/homework \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer TOKEN" \
  -d '{
    "title": "Geometry Assignment",
    "className": "Grade 10-A",
    "subject": "Mathematics",
    "description": "Complete chapter 7 exercises"
  }'

# 4️⃣ List your homework
curl -X GET http://localhost:8080/api/homeworks \
  -H "Authorization: Bearer TOKEN"

# 5️⃣ Check if server is healthy
curl http://localhost:8080/health
```

### Using Postman 📮
1. Import the OpenAPI spec from `pkg/openapi/openapi.yaml`
2. Set Authorization header: `Bearer <your-token>`
3. Start making requests! 🎉

---

## 🏗️ Architecture Philosophy (Why We Did It This Way)

### Problem ❌
- Monolithic code = spaghetti 🍝
- Tightly coupled = nightmare to test
- Framework locked-in = can't switch without rewriting everything
- Hard to understand = new devs cry 😭

### Solution ✅
- **Hexagonal Architecture** to the rescue!
- Business logic is independent (can test without HTTP/Database)
- Easy to swap implementations (MongoDB → PostgreSQL, Chi → Echo)
- Clean separation = happy developers 😊
- Future-proof design = your boss is proud 👔

### Benefits 🌟
```
Before: 🎡 Circular dependency nightmare
After:  ▶️  Clean unidirectional flow
        👑  Domain is king
        🔌  Easy to plug in new implementations
        🧪  Unit test everything
        📈  Scales with team growth
```

---

## 📝 Configuration (All the Switches)

### Server Configuration
```env
ENV=development          # development or production
HOST=localhost          # Server host
PORT=8080              # Server port
READ_TIMEOUT=10s       # Request read timeout
WRITE_TIMEOUT=10s      # Response write timeout
IDLE_TIMEOUT=120s      # Connection idle timeout
```

### Database Configuration
```env
MONGODB_URI=mongodb+srv://user:pass@cluster.mongodb.net/
DATABASE_NAME=classwork
DB_CONN_TIMEOUT=10s
DB_MAX_POOL_SIZE=100
DB_MIN_POOL_SIZE=10
```

### Authentication Configuration
```env
JWT_SECRET=your-super-secret-key-change-this-in-production
JWT_EXPIRATION=72h        # How long tokens last
JWT_REFRESH_TIME=24h      # When to refresh
```

---

## 🚀 Deployment (Ready for the World!)

### Docker 🐳
```bash
docker build -t classwork-api .
docker run -p 8080:8080 --env-file .env classwork-api
```

### Kubernetes ☸️
```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

### Traditional Server 🖥️
```bash
# Build
go build -o classwork .

# Copy binary and .env to server
scp classwork user@server:/app/
scp .env user@server:/app/

# Run (with process manager like systemd)
systemctl start classwork
```

---

## 📊 Monitoring & Logging 📈

The application logs to stdout with structured output. Integrate with:
- **ELK Stack** for centralized logging
- **Prometheus** for metrics
- **Grafana** for visualization
- **Datadog** for APM

---

## 🐛 Troubleshooting (When Things Go Wrong)

### MongoDB Connection Failed ❌
```
Check if MongoDB is running:
  - mongosh or mongo command
  - Verify MONGODB_URI in .env
  - Check firewall/network access
```

### JWT Token Issues 🔐
```
Make sure:
  - JWT_SECRET is set in .env
  - Authorization header format: "Bearer <token>"
  - Token hasn't expired
```

### Port Already in Use 🚫
```
Kill process using port 8080:
  Windows: netstat -ano | findstr :8080
  Linux/Mac: lsof -i :8080
  
Or change port:
  PORT=9000 go run .
```

---

## 🎓 Learning Resources

- 📖 [Hexagonal Architecture Guide](https://alistair.cockburn.us/hexagonal-architecture/)
- 🔐 [JWT Best Practices](https://tools.ietf.org/html/rfc7519)
- 🗄️  [MongoDB Documentation](https://docs.mongodb.com/)
- 🌐 [Chi Router Guide](https://github.com/go-chi/chi)
- 🔒 [OWASP Security Guide](https://owasp.org/)

---

## 📄 API Documentation

Full OpenAPI 3.0 documentation available at `pkg/openapi/openapi.yaml`

View it at: https://editor.swagger.io/ (paste your YAML file there)

---

## 🤝 Contributing

Found a bug? 🐛 Want to add a feature? 🚀

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

## 🙏 Acknowledgments

- The Go community for an amazing language ❤️
- MongoDB for reliable data storage 💾
- Chi Router for a clean HTTP implementation ⚡
- Coffee for making everything possible ☕

---

## 📞 Support

Need help? 
- 📧 Email: support@example.com
- 💬 Discussions: GitHub Issues
- 🐦 Twitter: @yourhandle

---

## 🎉 Final Thoughts

This project demonstrates:
- ✅ Clean architecture principles
- ✅ SOLID design patterns
- ✅ Modern Go practices
- ✅ Production-ready code
- ✅ Developer happiness 😊

Remember: **Bad code is temporary, good architecture is forever!**

---

**Built with ❤️ by developers who actually care about clean code**

_Last Updated: December 2025_
_Status: 🟢 Production Ready_

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
