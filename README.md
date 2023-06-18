# citi
Backend for Citi Bank Hackathon

# QR Code Generation from CNABs 750

This project is designed to process requests for generating QR codes from CNAB 750 files.

## Overview

The project consists of an API server application that handles HTTP requests for creating QR codes based on CNAB 750 files. It also includes a worker component responsible for periodically scanning a directory for new CNAB files and processing them.

## Getting Started

To run the project locally, follow these steps:

1. Clone the project repository: `git clone <repository-url>`
2. Navigate to the root directory of the project.
3. Install the dependencies by running `go mod download`.
4. Set the required environment variables:
   - `PORT`: The port on which the server will run.
5. Build and run the application using the command `go run cmd/main.go`.

## Project Structure

The project structure is as follows:

```
.
├── main.go
├── internal
│   ├── app
│   │   └── shipping_usecase.go
│   └── infra
│       ├── bacen_gateway.go
│       ├── return_storage.go
│       ├── shipping_handler.go
│       ├── shipping_repository.go
│       ├── shipping_worker.go
│       └── router.go
├── files
│   └── (CNAB 750 files will be stored here)
├── README.md
├── go.mod
└── go.sum
```

## Dependencies

The project uses the following external dependencies:

- `github.com/go-chi/chi` v1.5.4: Lightweight router for building HTTP APIs.
- `github.com/google/uuid` v1.3.0: Package for generating UUIDs.
- `github.com/julioc98/gocnab` v0.0.0-20230617184530-3abbcc8a6a90: CNAB file parser library.
- `github.com/go-chi/cors` v1.2.1: Middleware for handling Cross-Origin Resource Sharing (CORS).

Make sure to install these dependencies before running the application.

## Configuration

The application uses the following environment variables for configuration:

- `PORT`: The port on which the server will run.

Ensure that you have set this environment variable before running the application.

## Usage

Once the server is running, you can send HTTP requests to create QR codes based on CNAB 750 files. The server exposes the following endpoints:

- `POST /upload`: Upload a CNAB 750 file to generate QR codes.

## Worker

The worker component periodically scans the directory for new CNAB 750 files and processes them. It runs in the background and automatically detects and processes new files at regular intervals.

## Contributing

Contributions to the project are welcome. You can open issues for bug reports or submit pull requests for enhancements.

## License

This project is licensed under the [MIT License](LICENSE).
