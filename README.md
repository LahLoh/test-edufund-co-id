# TEST EDUFUND CO ID

## Features

### Login
Login is usecase for user login

#### Testing via cURL
```sh
curl http://localhost:50501/v1/login
```

#### Example Request & Response Login
```json
// Success when request match with database
// request
{
    "username": "admin@admin.com",
    "password": "adminadminadmin"
}

// response
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}

// Error when no payload or request body
// request


// response
{
    "error": "EOF"
}

// Error when no payload or request body username not valid
// request
{
    "username": "username",
    "password": "password"
}

// response
{
    "error": "please provide a valid email address"
}

// Error when no payload or request body password not valid
// request
{
    "username": "admin@admin.com",
    "password": "password"
}

// response
{
    "error": "password should be at least 12 characters long"
}

// Error when no payload or request body username or password not match with database
// request
{
    "username": "username",
    "password": "password"
}

// response
{
    "error": "invalid username / password"
}

// Error when no payload or request body fullfill but server error
// request
{
    "username": "admin@admin.com",
    "password": "adminadminadmin"
}

// response
{
    "error": "internal server error"
}
```

### Register
Register is usecase for user registration

#### Testing via cURL
```sh
curl http://localhost:50501/v1/register
```

#### Example Request & Response Register
```json
// Success when request match with database
// request
{
    "fullname":"admin",
    "username": "admin@admin.com",
    "password": "adminadminadmin",
    "confirmation_password": "adminadminadmin"
}

// response
{
    "message": "OK"
}

// Error when no payload or request body
// request


// response
{
    "error": "EOF"
}

// Error when no payload or request body fullname not valid
// request
{
    "fullname":"a",
    "username": "admin@admin.com",
    "password": "adminadminadmin",
    "confirmation_password": "adminadminadmin"
}

// response
{
    "error": "name should be 2 characters or more"
}

// Error when no payload or request body username not valid
// request
{
    "fullname":"admin",
    "username": "admin",
    "password": "adminadminadmin",
    "confirmation_password": "adminadminadmin"
}

// response
{
    "error": "please provide a valid email address"
}

// Error when no payload or request body password not valid
// request
{
    "fullname":"admin",
    "username": "admin",
    "password": "admin",
    "confirmation_password": "adminadminadmin"
}

// response
{
    "error": "password should be at least 12 characters long"
}

// Error when no payload or request body confirmation_password not valid
// request
{
    "fullname":"admin",
    "username": "admin",
    "password": "adminadminadmina",
    "confirmation_password": "adminadminadmin"
}

// response
{
    "error": "confirmation password does not match"
}

// Error when no payload or request body fullfill but server error
// request
{
    "fullname":"admin",
    "username": "admin@admin.com",
    "password": "adminadminadmin",
    "confirmation_password": "adminadminadmin"
}

// response
{
    "error": "internal server error"
}
```

## Tech Stack
- Docker (20.10.17) include `docker compose`
- Go (Go Language)
- MySQL (MySQL Database)
- goose (tool/lib for migration)

## How to Run with Docker Compose
```sh
docker compose up

# don't forget clean this image, if u want to re run, because it nge-cache
```

## How to Run without Docker Compose
```sh
# set env first
# DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_DATABASE, DB_OPTIONS, JWT_SECRET
go run main.go
```