# 🚀 Go Backend Boilerplate

A production-ready Go backend boilerplate built with **Gin**, **GORM**, **JWT Authentication**, **PostgreSQL**, **Redis**, **MongoDB**, **Rate Limiting**, and **Structured Logging**.

Designed using a clean layered architecture:

**Handler → Service → Repository → Database**

---

## ✨ Features

* REST API using Gin
* JWT Authentication & Authorization
* Role-Based Access Control (RBAC)
* PostgreSQL support via GORM
* MySQL support via GORM
* MongoDB support
* Redis integration
* Request Validation
* Structured Logging with Zap
* Rate Limiting
* CORS Support
* Graceful Shutdown
* Docker & Docker Compose Ready
* Clean Architecture

---

##  📂 Project Structure

```text
go-backend/
├── 📁 cmd/
│   └── 📁 server/
│       └── 📄 main.go
│
├── 📁 config/
│   ├── 📄 config.go
│   └── 📄 config.yaml
│
├── 📁 internal/
│   ├── 📁 database/
│   ├── 📁 models/
│   ├── 📁 repository/
│   ├── 📁 service/
│   ├── 📁 handler/
│   ├── 📁 middleware/
│   ├── 📁 router/
│   └── 📁 utils/
│
├── 📁 pkg/
│   └── 📁 logger/
│
├── 📄 .env
├── 📄 Dockerfile
├── 📄 docker-compose.yml
├── 📄 go.mod
└── 📄 go.sum
```

---

## 🧠 Architecture Overview

```text
           🌐 HTTP Request
                 │
                 ▼
        🧩 Handler Layer (Gin)
     ┌─────────────────────────┐
     │ - Parse Request         │
     │ - Validate Input        │
     │ - Call Service          │
     └──────────┬──────────────┘
                │
                ▼
        ⚙️ Service Layer
     ┌─────────────────────────┐
     │ - Business Logic        │
     │ - Rules & Validation    │
     │ - Transform Data        │
     └──────────┬──────────────┘
                │
                ▼
      🗄️ Repository Layer
     ┌─────────────────────────┐
     │ - Database Queries      │
     │ - CRUD Operations       │
     └──────────┬──────────────┘
                │
                ▼
        🗃️ Database Layer
     ┌─────────────────────────┐
     │ PostgreSQL / MySQL      │
     │ MongoDB / Redis         │
     └─────────────────────────┘
```

## 🧱 Layered Architecture Responsibilities

### 🧩 Handler Layer
Responsible for:
- 📥 Parsing incoming HTTP requests  
- ✅ Validating input data  
- 🔁 Calling service layer methods  
- 📤 Returning HTTP responses (JSON / status codes)  

---

### ⚙️ Service Layer
Responsible for:
- 🧠 Business logic implementation  
- 🔐 Authorization & role rules  
- 🔄 Data transformation between layers  
- 📊 Orchestrating repository calls  

---

### 🗄️ Repository Layer
Responsible for:
- 🧾 Database queries (CRUD operations)  
- 💾 Data persistence logic  
- 🔍 Fetching & filtering records  
- 🚫 No business logic (pure data access only)  

---

### 🗃️ Database Layer
Responsible for:
- 🐘 PostgreSQL (relational data)  
- 🐬 MySQL (relational data)  
- 🍃 MongoDB (document storage)  
- ⚡ Redis (cache / fast access data)  

---

# 🚀 Installation Guide

---

## 📥 Clone Repository

```bash
git clone https://github.com/kaustubh-dev/go-backend.git
cd go-backend
```

---

## 📦 Install Dependencies

### ⚡ Recommended (Automatic)

```bash
go mod tidy
```

### 🔍 What it does:
- 📥 Downloads missing dependencies  
- 🧹 Removes unused dependencies  
- 🔄 Updates `go.sum` automatically  

### 📌 Verify installed modules

```bash
go list -m all
```

---

## 🧰 Manual Dependency Installation

> Only use this if you want full control over dependencies

### 🌐 Web Framework

```bash
go get github.com/gin-gonic/gin@latest
```

### 🗄️ Database (ORM)

```bash
go get gorm.io/gorm@latest
go get gorm.io/driver/postgres@latest
go get gorm.io/driver/mysql@latest
```

### 🍃 MongoDB

```bash
go get go.mongodb.org/mongo-driver/mongo@latest
```

### ⚡ Redis

```bash
go get github.com/redis/go-redis/v9@latest
```

### 🔐 Authentication (JWT)

```bash
go get github.com/golang-jwt/jwt/v5@latest
```

### ⚙️ Configuration

```bash
go get github.com/spf13/viper@latest
go get github.com/joho/godotenv@latest
```

### ✅ Validation

```bash
go get github.com/go-playground/validator/v10@latest
```

### 📊 Logging

```bash
go get go.uber.org/zap@latest
```

### 🚦 Rate Limiting

```bash
go get golang.org/x/time/rate@latest
```

### 🆔 UUID

```bash
go get github.com/google/uuid@latest
```

### 🔒 Password Hashing

```bash
go get golang.org/x/crypto/bcrypt@latest
```

### 🌍 CORS

```bash
go get github.com/gin-contrib/cors@latest
```

---

# ⚙️ Environment Setup

Create a `.env` file in root directory:

```env
APP_ENV=development
APP_PORT=8080
APP_SECRET=super-secret-key

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=go_backend

MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=password
MYSQL_DB=go_backend

MONGO_URI=mongodb://localhost:27017
MONGO_DB=go_backend

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

JWT_SECRET=my-secret
JWT_EXPIRY_HOURS=24
JWT_REFRESH_EXPIRY_HOURS=168

RATE_LIMIT_RPS=5
RATE_LIMIT_BURST=10
```

---

# ▶️ Running the Application

## 🟢 Using Go

```bash
go run cmd/server/main.go
```

---

## 🛠️ Using Makefile

```makefile
run:
	go run cmd/server/main.go
```

Run:

```bash
make run
```

---

## 🐳 Using Docker

```bash
docker build -t go-backend .
docker run -p 8080:8080 go-backend
```

---

## 🧩 Using Docker Compose

```bash
docker-compose up --build
```

---

# 📡 API Endpoints

---

## 🔐 Authentication

| Method | Endpoint | Description |
|------|----------|-------------|
| POST | `/api/v1/auth/register` | Register user |
| POST | `/api/v1/auth/login` | Login user |

### 📌 Register User

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
-H "Content-Type: application/json" \
-d '{
  "name":"John Doe",
  "email":"john@example.com",
  "password":"password123"
}'
```

### 📌 Login User

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
-H "Content-Type: application/json" \
-d '{
  "email":"john@example.com",
  "password":"password123"
}'
```

---

## 👤 Users

| Method | Endpoint | Access |
|------|----------|--------|
| GET | `/api/v1/users/:id` | Authenticated |
| PATCH | `/api/v1/users/:id` | Authenticated |

### 📌 Get User

```bash
curl http://localhost:8080/api/v1/users/<USER_ID> \
-H "Authorization: Bearer <ACCESS_TOKEN>"
```

### 📌 Update User

```bash
curl -X PATCH http://localhost:8080/api/v1/users/<USER_ID> \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
-H "Content-Type: application/json" \
-d '{
  "name":"Updated Name"
}'
```

---

## 📦 Products

| Method | Endpoint | Access |
|------|----------|--------|
| GET | `/api/v1/products` | Authenticated |
| GET | `/api/v1/products/:id` | Authenticated |

### 📌 Get Products

```bash
curl http://localhost:8080/api/v1/products \
-H "Authorization: Bearer <ACCESS_TOKEN>"
```

### 📌 Get Product By ID

```bash
curl http://localhost:8080/api/v1/products/<PRODUCT_ID> \
-H "Authorization: Bearer <ACCESS_TOKEN>"
```

---

## 🛡️ Admin Routes

| Method | Endpoint |
|------|----------|
| GET | `/api/v1/admin/users` |
| DELETE | `/api/v1/admin/users/:id` |
| POST | `/api/v1/admin/products` |
| PUT | `/api/v1/admin/products/:id` |
| DELETE | `/api/v1/admin/products/:id` |

### 📌 Create Product

```bash
curl -X POST http://localhost:8080/api/v1/admin/products \
-H "Authorization: Bearer <ADMIN_TOKEN>" \
-H "Content-Type: application/json" \
-d '{
  "name":"iPhone 17",
  "description":"Latest Apple smartphone",
  "price":99999
}'
```

### 📌 Delete Product

```bash
curl -X DELETE http://localhost:8080/api/v1/admin/products/<PRODUCT_ID> \
-H "Authorization: Bearer <ADMIN_TOKEN>"
```

---

## ❤️ Health Check

| Method | Endpoint |
|------|----------|
| GET | `/health` |

```bash
curl http://localhost:8080/health
```

Response:

```json
{
  "status": "ok"
}
```

---

# 🔐 Authentication

```http
Authorization: Bearer <token>
```

---

# 🗄️ Auto Migration

```go
pgDB.AutoMigrate(
    &models.User{},
    &models.Product{},
)
```

---

# 📊 Logging

```json
{
  "level":"info",
  "msg":"HTTP Request",
  "method":"GET",
  "path":"/api/v1/users",
  "status":200
}
```

---

# 🚦 Rate Limiting

```env
RATE_LIMIT_RPS=5
RATE_LIMIT_BURST=10
```

---

# 🚀 Future Improvements

- 📘 Swagger Documentation  
- 🔄 Refresh Tokens  
- 📧 Email Verification  
- 🔑 Password Reset  
- 📈 Metrics & Monitoring  
- ☸️ Kubernetes Deployment  

---

# 🤝 Contributing

1. Fork repo  
2. Create feature branch  
3. Commit changes  
4. Push branch  
5. Create PR  

---

# 📜 License

This project is open-source and can be used as a production-ready Go backend template.
