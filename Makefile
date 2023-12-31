
APP_PKG = $(shell go list github.com/brcodingdev/chat-service/internal/...)

lint:
	@echo "Linting"
	@golint -set_exit_status $(APP_PKG)
	@golangci-lint run --timeout 3m0s

test:
	@echo "Testing"
	@go test ./... -v -count=1 -race

build:
	@echo "Building"
	@go build -o chatservice ./cmd

build-docker:
	@echo "Building docker image"
	@docker-compose build

run:
	@echo "Starting chat service"
	@docker-compose up -d
	@go run ./cmd

run-docker:
	@echo "Starting chat service with docker"
	@docker-compose up -d
