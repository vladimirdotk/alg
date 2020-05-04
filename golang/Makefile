.EXPORT_ALL_VARIABLES:

.PHONY: all
all: deps lint test

.PHONY: deps
deps:
	@go mod download
	@go mod vendor
	@go mod tidy

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test ./... -v