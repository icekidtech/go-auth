# Go Auth API

A simple authentication API built with Go, Fiber, PostgreSQL, and JWT. Features user signup, login, and protected endpoints with Swagger documentation.

## Features

- User registration (Signup)
- User login with JWT tokens
- Protected routes with JWT middleware
- Get current user profile
- Automatic password hashing with bcrypt
- PostgreSQL database with GORM ORM
- Swagger API documentation
- Comprehensive error handling

## Prerequisites

- **Go** 1.16+ ([Download](https://golang.org/dl/))
- **PostgreSQL** 12+ ([Download](https://www.postgresql.org/download/))
- **Git**

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/icekidtech/go-auth.git
cd go-auth
```

### 2. Install Dependencies

```bash
go mod download
go mod tidy
```

### 3. Create the Database

Open PostgreSQL and run:

```bash
psql -U postgres -f init.sql
```

Or manually:

```sql
CREATE DATABASE go_auth_db;
```

### 4. Configure Environment Variables

Copy `.env.example` to `.env` and update with your values:

```bash
cp .env.example .env
```

Edit `.env`:

```dotenv
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_actual_password
DB_NAME=go_auth_db

JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRY_HOURS=72

APP_PORT=3000
```

## Running the Application

```bash
go run main.go
```

You should see:

```
2026/05/17 16:00:00 Database connection established
2026/05/17 16:00:00 Database migrated successfully
2026/05/17 16:00:00 Server running on port 3000
```

## API Endpoints

### Health Check

```
GET /health
```

### Public Endpoints

#### Signup

```
POST /api/auth/signup
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (201):**

```json
{
  "token": "eyJhbGc...",
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2026-05-17T16:00:00Z"
  }
}
```

#### Login

```
POST /api/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (200):**

```json
{
  "token": "eyJhbGc...",
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2026-05-17T16:00:00Z"
  }
}
```

### Protected Endpoints

All protected endpoints require an `Authorization` header with a valid JWT token.

#### Get Current User

```
GET /api/user/me
Authorization: Bearer {token}
```

**Response (200):**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2026-05-17T16:00:00Z"
}
```

## Swagger API Documentation

Once the server is running, access the interactive Swagger documentation:

### Live Swagger UI

```
http://localhost:3000/swagger/index.html
```

Features:
- Browse all API endpoints
- Try out endpoints directly in the UI
- View request/response schemas
- Automatic API documentation

### Swagger JSON

```
http://localhost:3000/swagger/doc.json
```

## Project Structure

```
go-auth/
├── main.go              # Application entry point
├── go.mod              # Go module definition
├── .env.example        # Environment variables template
├── .gitignore          # Git ignore file
├── init.sql            # Database initialization script
├── README.md           # This file
│
├── config/
│   └── database.go     # Database configuration and connection
│
├── models/
│   └── user.go         # User model and request/response structs
│
├── handlers/
│   └── auth.go         # Authentication handlers (Signup, Login, Me)
│
├── middleware/
│   └── auth.go         # JWT authentication middleware
│
├── routes/
│   └── routes.go       # Route definitions
│
└── docs/               # Generated Swagger documentation
    ├── docs.go
    ├── swagger.json
    └── swagger.yaml
```

## Error Handling

The API returns appropriate HTTP status codes:

- `200 OK` - Successful request
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request format or missing required fields
- `401 Unauthorized` - Missing or invalid authentication token
- `409 Conflict` - Email already registered
- `500 Internal Server Error` - Server error

All errors include a JSON response with an error message:

```json
{
  "error": "Error message describing what went wrong"
}
```

## Security Considerations

1. **JWT Secret**: Change `JWT_SECRET` in `.env` to a strong, random value in production
2. **Password Hashing**: Passwords are hashed using bcrypt with default cost factor
3. **Token Expiry**: Tokens expire after 72 hours (configurable via `JWT_EXPIRY_HOURS`)
4. **HTTPS**: Use HTTPS in production
5. **Database**: Use strong database credentials and consider using connection pooling

## Development

### Adding New Endpoints

1. Add handler function in `handlers/auth.go`
2. Add Swagger documentation comments above the handler
3. Register the route in `routes/routes.go`
4. Run `swag init` to regenerate documentation

### Updating Swagger Docs

After making changes to handlers or API structure:

```bash
swag init
```

This regenerates the documentation in the `docs/` folder.

## Troubleshooting

### Database Connection Error

- Ensure PostgreSQL is running
- Verify database credentials in `.env`
- Check database exists: `psql -l`

### Port Already in Use

- Change `APP_PORT` in `.env` to an available port
- Or kill the process using port 3000

### JWT Token Invalid

- Ensure token is passed in Authorization header: `Authorization: Bearer <token>`
- Check token hasn't expired (default 72 hours)
- Verify `JWT_SECRET` is configured correctly

## License

MIT

## Support

For issues or questions, please open an issue on GitHub.
