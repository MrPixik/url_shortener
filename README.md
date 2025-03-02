# URL Shortener

A URL shortening service written in Go that allows users to submit URLs and receive shortened versions in return.

## Project Structure

```bash
url_shortener/
├── cmd/                            #
│   ├── client/                     # Client-side implementation
│   │   └── main.go                 #
│   │   └── clients.go              #
│   │   └── models.go               #
│   │   └── menu.go                 #
│   ├── shortener/                  # Server-side implementation
│   │   ├── main.go                 # Main entry point for the shortener service
├── internal/                       #
│   ├── app/                        #
│   │   ├── middleware/             # Middleware logic
│   │   │   ├── compressing.go      #     
│   │   │   ├── logging.go          #
│   │   ├── models/                 # Models
│   │   │   ├── easyjson/           # EasyJSON files
│   │   │   │   ├── server.go       # Models for requests and responses processing
│   │   │   ├── db.go               # Models for working with database client 
│   │   │   ├── url_file.go         # Models for working with file ///OUTDATED///
│   │   ├── server/                 # Server handlers and tests
│   │   │   ├── handler.go          # URL shortening handler
│   │   │   ├── handler_test.go     # Unit tests for the handler
│   └── config/                     # Configuration
│   │   └── config.go
│   ├── db/                         # Database logic
│   │   ├── mocks/                  # Mocks for unit-testing
│   │   │   ├── mocks_db_service.go # 
│   │   ├── init.go                 # DatabaseService initialization
│   │   ├── interfaces.go           #
│   │   ├── models.go               #
├── go.mod                          # Go module file
└── .gitignore                      # Git ignore file
```
## Installation

To install and run this project locally:

1. Clone the repository:

``` bash
git clone https://github.com/MrPixik/url_shortener.git
cd url_shortener
```
2. Install dependencies:

```bash
go mod tidy
```
3. Build and run the server:

```bash
go run cmd/shortener/main.go
```
4. Now you can interact with the server, using client menu:

```bash
go run cmd/client/main.go
```
Also you can use any HTTP client (e.g., curl or Postman).

## API Documentation

| HTTP Method | Endpoint                          | Description                                              |
|-------------|-----------------------------------|----------------------------------------------------------|
| **POST**    | `/`                               | Create a new short URL. (Content-Type: text/plain)       |
| **POST**    | `/api/shorten`                    | Create a new short URL. (Content-Type: application/json) |
| **POST**    | `/api/shorten/batch`              | Create a new short URLs.                                 |
| **GET**     | `/{id}`                           | Go to URL, using it's short implementation.              |
| **GET**     | `/ping`                           | Ping to database.                                        |


## Configuration

The server can be configured through the internal/config/config.go file, where you can specify parameters such as:

ShortURLAddr - the base address for shortened URLs

## Testing

Unit tests are located in internal/app/server/handler_test.go. To run the tests, use the following command:

```bash
go test ./...
```
## Middleware

Logging: The project includes a logging middleware in internal/app/middleware/logging.go that logs incoming requests and responses.

Compressing: The project includes a logging middleware in internal/app/middleware/compressing.go that provides data compression and unpacking for incoming requests and responses. 

## JSON Handling with EasyJSON

The project uses EasyJSON for high-performance JSON serialization and deserialization. Models for the request and response are located in internal/app/models/easy_json_models.go. The EasyJSON code is automatically generated using the go generate command.

To regenerate EasyJSON files, run:

```bash
go generate ./...
```
## CI/CD
This project includes GitHub Actions workflows for continuous integration:

shortenertest.yml: Runs the tests on each push.
staticcheck.yml: Performs static analysis using Go linters.
