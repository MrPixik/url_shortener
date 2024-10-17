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
│   ├── shortener/               # Server-side implementation
│   │   ├── main.go              # Main entry point for the shortener service
│   │   └── README.md            # Documentation for the shortener
├── internal/
│   ├── app/
│   │   ├── middleware/          # Middleware logic (e.g., logging)
│   │   ├── models/              # Models and EasyJSON files
│   │   │   ├── easy_json_models.go         # Models for URLRequest and URLResponse
│   │   │   └── easy_json_models_easyjson.go # Generated EasyJSON code
│   │   ├── server/              # Server handlers and tests
│   │   │   ├── handler.go       # URL shortening handler
│   │   │   ├── handler_test.go  # Unit tests for the handler
│   └── config/                  # Configuration
│       └── config.go
├── go.mod                       # Go module file
└── .gitignore                   # Git ignore file
