# Notes API

## Overview

This is a simple API for managing notes, implemented in Go (Golang) using the [Gin](https://github.com/gin-gonic/gin) framework and PostgreSQL as the database.

## Table of Contents

- [Architecture](#architecture)
- [Framework](#Framework)
- [Database](#Database)
- [Third-Party-Tools](#Third-Party-Tools)
- [JWT-Authentication](#JWT-Authentication)
- [Setup](#setup)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)

## Architecture

The application is structured in a standard MVC (Model-View-Controller) architecture.

- **Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL
- **Authentication**: JWT tokens
- **Validation**: [Go Playground Validator](https://github.com/go-playground/validator)
- **ORM**: [GORM](https://gorm.io/)

## Framework
### Gin
Gin is a web framework written in Go (Golang). It features a fast router, middlewares support, and a robust set of tools for building web applications. Gin was chosen for its simplicity, performance, and ease of use in building RESTful APIs.

## Database
### PostgreSQL
PostgreSQL is a powerful, open-source relational database system. It was chosen for its reliability, support for complex queries, and the ability to handle large amounts of data. The database is used to store user information, notes, and shared notes.

## Third-Party-Tools
### GORM
GORM is an Object Relational Mapping (ORM) library for Golang. It simplifies database operations and provides a convenient way to interact with the database. GORM is used for database migrations, creating, and querying records.

### Validator
Validator is a powerful validator library for Go. It is used for validating the request payloads and ensuring that the data meets the required constraints.

## JWT-Authentication
JWT is a standard for creating tokens that assert some number of claims. This API uses JWT for secure user authentication.

## Setup

1. **Clone the repository:**

    ```bash
    git clone https://github.com/prathishbv/notes-api.git
    cd notes-api
    ```

2. **Install dependencies:**

    ```bash
    go mod install
    go mod tidy
    ```

3. **Database Setup:**

    - Create a PostgreSQL database and update the configuration in `.env`.

4. **Run the application:**

    ```bash
    go run main.go
    ```

## Usage

- The API will be accessible at `http://localhost:PORT` where `PORT` is the specified port in your configuration.

- Use tools like [Postman](https://www.postman.com/) or [curl](https://curl.se/) to interact with the API.

## API Endpoints

- **Authentication:**
    - `POST /api/auth/signup`: Register a new user.
    - `POST /api/auth/login`: Log in and get a JWT token.

- **User Management:**
    - `GET /api/users`: Get all users.

- **Notes Management:**
    - `GET /api/notes`: Get all notes.
    - `GET /api/notes/:id`: Get a specific note.
    - `POST /api/notes`: Create a new note.
    - `PUT /api/notes/:id`: Update a note.
    - `DELETE /api/notes/:id`: Delete a note.
    - `POST /api/notes/:id/share`: Share a note with another user.
    - `GET /api/search`: Search for notes based on a query.


