.PHONY: build run dev test fmt lint clean install-air

# Build the Go binary
build:
	go build -o bin/app main.go

# Run the app
run:
	go run main.go

# Run with live-reload using Air
dev:
	~/go/bin/air -c .air.toml

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Remove built binary
clean:
	rm -rf tmp

# Install Air for live-reload
install-air:
	go install github.com/cosmtrek/air@latest
