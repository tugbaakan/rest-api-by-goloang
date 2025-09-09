# GO REST API Example

This is a RESTful API built with Go, Gin, and SQLite. It supports user registration, authentication (JWT), event management, and event registration.

## Features

- User sign-up and login with password hashing
- JWT-based authentication middleware
- CRUD operations for events
- Register/cancel registration for events
- SQLite database (auto-creates tables on startup)
- Example HTTP requests in `api-tests/`

## Project Structure

- `main.go` – Entry point
- `db/` – Database initialization
- `models/` – Data models (`User`, `Event`)
- `routes/` – HTTP route handlers
- `middlewares/` – Authentication middleware
- `utils/` – Utility functions (JWT, password hashing)
- `api-tests/` – Example HTTP requests

## Getting Started

### Prerequisites

- Go 1.25+
- SQLite

### Install Dependencies

```sh
go mod tidy
```

### Run the Server

```sh
go run main.go
```

Server runs on `localhost:8080`.

### API Endpoints

- `POST /signup` – Register user
- `POST /login` – Login, returns JWT token
- `GET /events` – List events
- `GET /events/:id` – Get event details
- `POST /events` – Create event (auth required)
- `PUT /events/:id` – Update event (auth required)
- `DELETE /events/:id` – Delete event (auth required)
- `POST /events/:id/register` – Register for event (auth required)
- `DELETE /events/:id/register` – Cancel registration (auth required)

### Example Requests

See [api-tests/](api-tests/) for sample HTTP requests.

##