# API Testing Guide

Quick reference for testing all API endpoints.

## Base URL
```
http://localhost:8080
```

## 1. Health Check

**Endpoint**: `GET /health`

**Description**: Check if the server is running

**Request**:
```bash
curl http://localhost:8080/health
```

**Response** (200 OK):
```json
{
  "status": "healthy"
}
```

---

## 2. Register Teacher

**Endpoint**: `POST /api/register`

**Description**: Create a new teacher account

**Request**:
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "Password123!"
  }'
```

**Request Body**:
```json
{
  "name": "string (required)",
  "email": "string (email format, required)",
  "password": "string (min 6 chars, required)"
}
```

**Response** (201 Created):
```json
{
  "id": "507f1f77bcf86cd799439011"
}
```

**Response** (400 Bad Request):
```json
{
  "error": "email already taken"
}
```

**Validation Rules**:
- Name: Required
- Email: Required, must be valid email format
- Password: Required, minimum 6 characters

---

## 3. Login Teacher

**Endpoint**: `POST /api/login`

**Description**: Authenticate and get JWT token

**Request**:
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "Password123!"
  }'
```

**Request Body**:
```json
{
  "email": "string (email format, required)",
  "password": "string (required)"
}
```

**Response** (200 OK):
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZWFjaGVyX2lkIjoiNTA3ZjFmNzdiY2Y4NmNkNzk5NDM5MDExIiwiZXhwIjoxNzMzNzQ2NjkwfQ...."
}
```

**Response** (401 Unauthorized):
```json
{
  "error": "invalid credentials"
}
```

**Note**: Save the token for authenticated requests. The token is valid for 72 hours by default.

---

## 4. Create Homework

**Endpoint**: `POST /api/homework`

**Required**: Bearer token (from login)

**Description**: Create a new homework assignment

**Request**:
```bash
TOKEN="your_token_from_login"

curl -X POST http://localhost:8080/api/homework \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "Math Assignment Chapter 1",
    "description": "Complete exercises 1-10 from chapter 1",
    "className": "Class 10-A",
    "subject": "Mathematics",
    "attachments": "https://example.com/chapter1.pdf"
  }'
```

**Request Body**:
```json
{
  "title": "string (required)",
  "description": "string (optional)",
  "className": "string (required)",
  "subject": "string (required)",
  "attachments": "string (optional, URL or file reference)"
}
```

**Response** (201 Created):
```json
{
  "id": "507f1f77bcf86cd799439012",
  "title": "Math Assignment Chapter 1",
  "description": "Complete exercises 1-10 from chapter 1",
  "className": "Class 10-A",
  "subject": "Mathematics",
  "teacherId": "507f1f77bcf86cd799439011",
  "attachments": "https://example.com/chapter1.pdf",
  "createdAt": "2024-12-08T15:30:00Z",
  "updatedAt": "2024-12-08T15:30:00Z"
}
```

**Response** (401 Unauthorized):
```json
{
  "error": "invalid token"
}
```

---

## 5. Get All Homeworks

**Endpoint**: `GET /api/homeworks`

**Required**: Bearer token (from login)

**Description**: List all homework assignments for the authenticated teacher

**Request**:
```bash
TOKEN="your_token_from_login"

curl -X GET http://localhost:8080/api/homeworks \
  -H "Authorization: Bearer $TOKEN"
```

**Response** (200 OK):
```json
[
  {
    "id": "507f1f77bcf86cd799439012",
    "title": "Math Assignment Chapter 1",
    "description": "Complete exercises 1-10 from chapter 1",
    "className": "Class 10-A",
    "subject": "Mathematics",
    "teacherId": "507f1f77bcf86cd799439011",
    "attachments": "https://example.com/chapter1.pdf",
    "createdAt": "2024-12-08T15:30:00Z",
    "updatedAt": "2024-12-08T15:30:00Z"
  },
  {
    "id": "507f1f77bcf86cd799439013",
    "title": "English Essay",
    "description": "Write an essay on Shakespeare",
    "className": "Class 10-A",
    "subject": "English",
    "teacherId": "507f1f77bcf86cd799439011",
    "attachments": "https://example.com/essay-guidelines.pdf",
    "createdAt": "2024-12-08T16:00:00Z",
    "updatedAt": "2024-12-08T16:00:00Z"
  }
]
```

**Response** (empty list):
```json
[]
```

**Response** (401 Unauthorized):
```json
{
  "error": "invalid token"
}
```

---

## 6. Get Single Homework

**Endpoint**: `GET /api/homework/:id`

**Required**: Bearer token (from login)

**Description**: Get details of a specific homework assignment

**Request**:
```bash
TOKEN="your_token_from_login"
HOMEWORK_ID="507f1f77bcf86cd799439012"

curl -X GET http://localhost:8080/api/homework/$HOMEWORK_ID \
  -H "Authorization: Bearer $TOKEN"
```

**Response** (200 OK):
```json
{
  "id": "507f1f77bcf86cd799439012",
  "title": "Math Assignment Chapter 1",
  "description": "Complete exercises 1-10 from chapter 1",
  "className": "Class 10-A",
  "subject": "Mathematics",
  "teacherId": "507f1f77bcf86cd799439011",
  "attachments": "https://example.com/chapter1.pdf",
  "createdAt": "2024-12-08T15:30:00Z",
  "updatedAt": "2024-12-08T15:30:00Z"
}
```

**Response** (404 Not Found):
```json
{
  "error": "homework not found"
}
```

**Response** (400 Bad Request - Invalid ID):
```json
{
  "error": "invalid homework_id"
}
```

---

## 7. Update Homework

**Endpoint**: `PUT /api/homework/:id`

**Required**: Bearer token (from login)

**Description**: Update an existing homework assignment

**Request**:
```bash
TOKEN="your_token_from_login"
HOMEWORK_ID="507f1f77bcf86cd799439012"

curl -X PUT http://localhost:8080/api/homework/$HOMEWORK_ID \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "Math Assignment Chapter 1-2 (Updated)",
    "description": "Complete exercises 1-15 from chapters 1-2"
  }'
```

**Request Body** (all fields optional):
```json
{
  "title": "string (optional)",
  "description": "string (optional)",
  "className": "string (optional)",
  "subject": "string (optional)",
  "attachments": "string (optional)"
}
```

**Response** (200 OK):
```json
{
  "id": "507f1f77bcf86cd799439012",
  "title": "Math Assignment Chapter 1-2 (Updated)",
  "description": "Complete exercises 1-15 from chapters 1-2",
  "className": "Class 10-A",
  "subject": "Mathematics",
  "teacherId": "507f1f77bcf86cd799439011",
  "attachments": "https://example.com/chapter1.pdf",
  "createdAt": "2024-12-08T15:30:00Z",
  "updatedAt": "2024-12-08T16:15:00Z"
}
```

**Response** (404 Not Found):
```json
{
  "error": "homework not found"
}
```

---

## 8. Delete Homework

**Endpoint**: `DELETE /api/homework/:id`

**Required**: Bearer token (from login)

**Description**: Delete a homework assignment

**Request**:
```bash
TOKEN="your_token_from_login"
HOMEWORK_ID="507f1f77bcf86cd799439012"

curl -X DELETE http://localhost:8080/api/homework/$HOMEWORK_ID \
  -H "Authorization: Bearer $TOKEN"
```

**Response** (200 OK):
```json
{
  "message": "homework deleted successfully"
}
```

**Response** (404 Not Found):
```json
{
  "error": "homework not found"
}
```

---

## Complete Example Workflow

### Step 1: Register
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "email": "jane@example.com",
    "password": "SecurePass123"
  }'
# Save the ID from response
```

### Step 2: Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "jane@example.com",
    "password": "SecurePass123"
  }'
# Save the token from response
```

### Step 3: Create Homework
```bash
TOKEN="your_token_here"

curl -X POST http://localhost:8080/api/homework \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "Physics Lab Report",
    "description": "Complete the lab on Newton's Laws",
    "className": "Class 11-B",
    "subject": "Physics",
    "attachments": "https://example.com/lab-instructions.pdf"
  }'
# Save the homework ID from response
```

### Step 4: List Homeworks
```bash
curl -X GET http://localhost:8080/api/homeworks \
  -H "Authorization: Bearer $TOKEN"
```

### Step 5: Update Homework
```bash
HOMEWORK_ID="homework_id_from_step_3"

curl -X PUT http://localhost:8080/api/homework/$HOMEWORK_ID \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "description": "Complete the lab on Newton's Laws (due by Friday)"
  }'
```

### Step 6: Delete Homework
```bash
curl -X DELETE http://localhost:8080/api/homework/$HOMEWORK_ID \
  -H "Authorization: Bearer $TOKEN"
```

---

## Error Handling

### Common HTTP Status Codes

| Code | Meaning | Example |
|------|---------|---------|
| 200 | OK | Successful GET/PUT request |
| 201 | Created | Successful POST request |
| 400 | Bad Request | Invalid input data |
| 401 | Unauthorized | Missing or invalid token |
| 404 | Not Found | Resource doesn't exist |
| 500 | Server Error | Unexpected server error |

### Common Error Messages

| Error | Cause | Solution |
|-------|-------|----------|
| "invalid token" | Token expired or malformed | Login again to get new token |
| "email already taken" | Email already registered | Use different email |
| "invalid credentials" | Wrong email/password | Check email and password |
| "homework not found" | Invalid homework ID | Verify the homework ID |
| "missing authorization header" | No Bearer token sent | Add Authorization header |

---

## Tools for Testing

### Using Postman
1. Open Postman
2. Create new HTTP request
3. Copy request examples above
4. Set headers (Content-Type: application/json, Authorization: Bearer TOKEN)
5. Send request
6. View response

### Using curl (Command Line)
Shown in examples above

### Using VS Code REST Client Extension
Install "REST Client" extension, create `.http` file:
```http
### Register
POST http://localhost:8080/api/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "Password123!"
}

### Login
POST http://localhost:8080/api/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "Password123!"
}
```

---

## Authentication Header Format

All protected endpoints require this header:
```
Authorization: Bearer <your_jwt_token_here>
```

Example:
```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  http://localhost:8080/api/homeworks
```

---

## Rate Limiting & Timeouts

- **Request Timeout**: 10 seconds (configurable)
- **No rate limiting**: Currently unlimited (can be added)
- **Token Expiration**: 72 hours (configurable)

---

For more information, see:
- README.md - Project overview
- ARCHITECTURE.md - Technical architecture
- DEVELOPMENT.md - Development guide
