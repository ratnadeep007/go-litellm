# AGENTS.md

## Build, Lint, and Test Commands

### Build
- To run the server:
  ```bash
  make server
  ```

### Lint
- Use `golangci-lint` for linting:
  ```bash
  golangci-lint run
  ```

### Test
- Run all tests:
  ```bash
  go test ./...
  ```
- Run a specific test:
  ```bash
  go test -run TestName
  ```

## Code Style Guidelines

### Imports
- Group imports into three sections: standard library, third-party libraries, and local packages.
- Use aliases for imports only when necessary.

### Formatting
- Use `gofmt` to format code:
  ```bash
  gofmt -s -w .
  ```

### Types
- Use descriptive names for types and variables.
- Avoid single-letter names except for loop indices.

### Naming Conventions
- Use `PascalCase` for exported functions and types.
- Use `camelCase` for unexported functions and variables.

### Error Handling
- Always check for errors and handle them appropriately.
- Use `fmt.Errorf` to wrap errors with context.

### Environment Variables
- Use the `checkEnvVar` function in `lib/utils.go` to validate required environment variables.

### Logging
- Use the `echo` middleware for logging HTTP requests and errors.

### API Endpoints
- Define endpoints in `server/main.go` and implement handlers in `server/handlers`.

### Supported Providers
- This project supports routing requests to multiple LLM providers and diffusion models, such as:
  - OpenAI
  - Stability AI
  - Other compatible APIs

---

This file provides essential guidelines and commands for agents working in this repository.