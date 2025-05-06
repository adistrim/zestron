# Zestron Server

Go backend based on Gin framework.

## Setup
1. Install Go (1.24+).
2. Create `.env` with `KEY`.
3. Run `go mod tidy`.
4. Build: `go build -o dist/zestron-server main.go`.
5. Run: `GIN_MODE=release ./dist/zestron-server`.
6. Test: `curl -X POST http://localhost:8080/api/generate -H "Content-Type: application/json" -d '{"prompt":"What is the capital of India?"}'`.

## Endpoint
- `GET /ping`: Check if server is running.
- `POST /api/generate`: Send prompt to the LLM API, returns response.
