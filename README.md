# Docker Example Project - Go Debug Server

A simple Go webserver that responds with helpful debug information in JSON format.

## Features

- Listens on port 8080
- Returns JSON with debug information including:
  - Request timestamp
  - HTTP method and path
  - Client remote address
  - Request headers and query parameters
  - Server information (hostname, Go version, CPU count, OS, architecture)

## Running Locally

```bash
go run main.go
```

Then visit `http://localhost:8080` or use curl:

```bash
curl http://localhost:8080
```

## Docker Build & Run

Build the Docker image:

```bash
docker build -t debug-server .
```

Run the container:

```bash
docker run -p 8080:8080 debug-server
```

## Example Response

```json
{
  "timestamp": "2025-10-23T08:57:16Z",
  "method": "GET",
  "path": "/",
  "remote_addr": "172.17.0.1:54321",
  "headers": {
    "User-Agent": "curl/7.68.0",
    "Accept": "*/*"
  },
  "query_params": {},
  "host": "localhost:8080",
  "server_info": {
    "hostname": "abc123def456",
    "go_version": "go1.25",
    "num_cpu": 8,
    "num_goroutine": 3,
    "os": "linux",
    "architecture": "amd64"
  }
}
```
