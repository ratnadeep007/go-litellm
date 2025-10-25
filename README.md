# Go LiteLLM

A lightweight Go-based server for generating responses using OpenAI's API.

## Prerequisites

- Go 1.20 or later installed on your system.
- An OpenAI API key. You can obtain one from [OpenAI](https://platform.openai.com/).

## Getting Started

Follow these steps to set up and run the server:

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/golang-litellm.git
cd golang-litellm
```

### 2. Set Up Environment Variables

Export your OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY=your_openai_api_key
```

### 3. Install Dependencies

Ensure all dependencies are installed:

```bash
go mod tidy
```

### 4. Run the Server

Start the server using the `make` command:

```bash
make server
```

The server will start on `http://localhost:1323`.

## API Endpoints

### `POST /generate`

Generate a response using OpenAI's API. Send a JSON payload with the required input.

Example:

```bash
curl -X POST http://localhost:1323/generate \
-H "Content-Type: application/json" \
-d '{"prompt": "Hello, how are you?"}'
```

## Development

### Linting

Run the linter to ensure code quality:

```bash
golangci-lint run
```

### Testing

Run all tests:

```bash
go test ./...
```

Run a specific test:

```bash
go test -run TestName
```

### Formatting

Format the codebase:

```bash
gofmt -s -w .
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.