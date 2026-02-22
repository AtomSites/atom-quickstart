.PHONY: install dev build clean test lint

install:
	@echo "Installing Go dependencies..."
	go mod download
	@echo "Installing development tools..."
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/cespare/reflex@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Installation complete!"

check-templ:
	@command -v templ >/dev/null 2>&1 || { echo "templ not found. Run: make install" >&2; exit 1; }

check-reflex:
	@command -v reflex >/dev/null 2>&1 || { echo "reflex not found. Run: make install" >&2; exit 1; }

dev: check-templ check-reflex
	@echo "Starting development server with hot reload..."
	@templ generate
	reflex -r '\.(go|templ)$$' -R '_templ\.go$$' -s -- sh -c 'templ generate && go run ./cmd/server'

build: check-templ
	@echo "Building binary..."
	templ generate
	go build -o bin/server ./cmd/server

test: check-templ
	@echo "Running tests..."
	templ generate
	go test ./... -v

lint:
	@echo "Running linter..."
	golangci-lint run ./...

clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
