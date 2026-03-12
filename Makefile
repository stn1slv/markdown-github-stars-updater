VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: setup test lint format build run upgrade-deps

# Bootstrap the project
setup:
	go mod tidy
	go mod download

# Run unit tests with coverage
test:
	go test -race -coverprofile=coverage.out ./...

# Run linters and static analysis using pinned tools
lint:
	go tool golangci-lint run
	go tool govulncheck ./...
	go tool nilaway ./...

# Auto-format code
format:
	go tool gofumpt -extra -w .

# Compile the application for the current platform
build:
	CGO_ENABLED=0 GOFLAGS=-trimpath go build -ldflags="$(LDFLAGS)" -o updater

# Run the application (example usage)
run:
	go run main.go

# Upgrade all dependencies to latest versions
upgrade-deps:
	go get -u ./...
	go mod tidy
