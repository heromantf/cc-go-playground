# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Run Commands

```bash
# Run the application
go run main.go

# Build the application
go build -o cc-go-playground

# Run tests
go test ./...

# Run a single test
go test -run TestName ./...
```

## Project Overview

This is a Go playground/sandbox project for experimentation. It uses Go 1.25 and implements a demo REST API server using the Gin web framework.

## API Endpoints

The server runs on port 8080 and provides the following endpoints:

- `GET /` - Welcome message
- `GET /health` - Health check endpoint
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create a new user (requires `name` and `email` in JSON body)
- `PUT /api/v1/users/:id` - Update user by ID
- `DELETE /api/v1/users/:id` - Delete user by ID
