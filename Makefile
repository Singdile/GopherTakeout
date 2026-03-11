.PHONY: run build test test-unit test-integration test-cover lint fmt migrate-up migrate-down db-reset clean help

APP_NAME=GopherTakeout
MAIN_PATH=./cmd/server
BUILD_DIR=./bin

GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOFMT=gofmt
GOLINT=golangci-lint

.DEFAULT_GOAL := help

run:
	$(GORUN) $(MAIN_PATH)/main.go

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

test-unit:
	$(GOTEST) -v -short ./...

test-integration:
	$(GOTEST) -v ./test/integration/...

test-cover:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

lint:
	$(GOLINT) run ./...

fmt:
	$(GOFMT) -s -w .

migrate-up:
	@echo "Run: psql -U postgres -d gopher_takeout -f migrations/001_init.sql"

migrate-down:
	@echo "Please rollback migrations manually"

db-reset:
	@echo "Please reset database manually"

clean:
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

deps:
	$(GOCMD) mod download
	$(GOCMD) mod tidy

help:
	@echo "GopherTakeout - Makefile Commands"
	@echo ""
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'