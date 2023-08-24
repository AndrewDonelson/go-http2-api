# go-http2-api


This repository contains an HTTP/2 API server implementation in Go. The server is designed with a focus on Quick Start.

## Project Structure

- `api.go`: Contains the core logic for the API server setup, including structures for API routes and the server itself.
- `handlers.go`: Contains public handler functions, including a basic heartbeat handler.
- `helpers.go`: Contains utility functions related to the API server.
- `main.go`: Main entry point for starting the server.

## Installation and Setup

1. Clone the repository:
```bash
git clone https://github.com/AndrewDonelson/go-http2-api
cd go-http2-api
```

2. Install the required Go dependencies:
```bash
go mod tidy
```

## Running the Project

Navigate to the root directory and run:

```bash
cd cmd/apisvr
go run main.go
```

The server will start on port 8080, offering HTTP/2 support.

## API Endpoints Overview

The server provides several endpoints for zero-knowledge proofs and data broadcast:

- `/heartbeat`: Basic endpoint to check if the service is alive.
- ... (Add other endpoints as needed)

## Contribution

Feel free to submit pull requests or raise issues.

## License

[MIT License](LICENSE)