.PHONY: build run test lint tidy fmt vet

APP_NAME := api
BUILD_DIR := bin
MAIN_PKG := ./cmd/api

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PKG)

run:
	go run $(MAIN_PKG)

test:
	go test ./... -race -count=1

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed, falling back to 'go vet'"; \
		go vet ./...; \
	fi

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...
