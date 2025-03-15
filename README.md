# Go JWT Example - Rest API

A simple Go-based REST API project using JWT authentication, GORM for ORM, and Gin as the web framework. This project demonstrates how to set up a secure API with user authentication, posts management, database migrations, and API versioning.

## Features

- **JWT Authentication** for secure endpoints
- **Versioned API endpoints**
- **Database migration** with GORM and Gormigrate
- **Seeding of initial data** (with conditional checks to avoid re-seeding)
- **Environment variable configuration** for flexibility

## Requirements

- Go 1.18+
- MySQL (or any other database that GORM supports)
- GIN web framework
- GORM for ORM
- Gormigrate for database migrations

## Setup Instructions

### Clone the Repository

```bash
git clone https://github.com/jobayer-mojumder/go-jwt-example.git
cd go-jwt-example
```

### Install Dependencies

Use Go modules to install the necessary dependencies:

```bash
go mod tidy
```

### Configure Environment Variables

Create a `.env` file in the root directory of the project and set your environment variables. Example:

```ini
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourdbuser
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
JWT_SECRET=yourjwtsecret
```

### Database Setup

Ensure your database is up and running. I am using MySQL for this project. Update the `.env` file with the correct credentials and configure your database accordingly.

### Run the Application

Migrations and seeders will automatically run when the project starts. To start the server, use the following command:

```bash
go run ./cmd/api
```

### API Endpoints

The following endpoints are available:

#### Public Routes
- `GET /v1/`: Returns a "Hello World" message.
- `POST /v1/login`: Logs in a user and returns a JWT token.

#### Private Routes (Requires JWT token in Authorization header)
- `GET /v1/posts`: Returns a list of posts.
- `POST /v1/posts`: Creates a new post.

You can test the API using tools like Postman or `curl`.

### Testing Authentication

To test authentication, you need to log in via `POST /v1/login` with a username and password (hardcoded in the seed). The response will contain a JWT token. Pass this token in the `Authorization` header as `Bearer <your_token>` for any subsequent requests to protected routes.


## Database Migrations

This project uses **Gormigrate** for database migrations.

To create new migrations, create a migration file in the `migrations` folder and run the migrations in `main.go`. For example:

```go
package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreatePostTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2025_03_15_002", // Migration ID
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Post{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&models.Post{})
		},
	}
}
```