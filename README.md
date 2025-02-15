# URL Shortener

A URL shortening service written in Go that allows users to submit URLs and receive shortened versions in return.

## Project Structure

```bash
url_shortener/
├── .github/                    # GitHub workflows for CI/CD
│   ├── workflows/
│   │   ├── shortenertest.yml    # Test workflow
│   │   └── statictest.yml       # Linter/static analysis workflow
├── cmd/
│   ├── client/                  # Client-side implementation
│   │   └── main.go
│   │   └── clients.go
│   │   └── models.go
│   │   └── menu.go
│   ├── shortener/               # Server-side implementation
│   │   ├── main.go              # Main entry point for the shortener service
│   │   └── README.md            # Documentation for the shortener
├── internal/
│   ├── app/
│   │   ├── middleware/          # Middleware logic (e.g., logging, compressing# )
│   │   ├── models/              # Models and EasyJSON files
│   │   │   ├── easy_json_models.go         # Models for URLRequest and URLResponse
│   │   │   └── easy_json_models_easyjson.go # Generated EasyJSON code
│   │   ├── server/              # Server handlers and tests
│   │   │   ├── handler.go       # URL shortening handler
│   │   │   ├── handler_test.go  # Unit tests for the handler
│   └── config/                  # Configuration
│   │   └── config.go
│   ├── db/                      # Database logic
│   │   ├── init.go              # DatabaseService initialization
│   │   ├── interfaces.go   
│   │   ├── models.go   
├── tmp/                         # Temp files
├── go.mod                       # Go module file
└── .gitignore                   # Git ignore file
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
cd cmd/shortener
go run main.go
```
4. You can now interact with the server using any HTTP client (e.g., curl or Postman).

## Usage

Shortening a URL
To shorten a URL, send a POST request to the server with a JSON payload containing the URL to be shortened:

```bash
curl -X POST http://localhost:8080/shorten \
     -H "Content-Type: application/json" \
     -d '{"url":"https://www.example.com"}'
```
Example response:

```json
{
  "url": "http://localhost:8080/abc123"
}
```
Expanding a URL
To retrieve the original URL, send a GET request to the shortened URL:

```bash
curl http://localhost:8080/abc123
```
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
