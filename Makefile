.PHONY: help
help:  ## Show this helper
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.PHONY: build-dev
build-dev:  ## Build the application for development purposes
	@go build -o passager cmd/passager/main.go

.PHONY: format
format:  ## Apply go fmt to all files
	@go fmt ./...

.PHONY: tests
tests:  ## Run all unit tests and show the code coverage report
	@go test -coverprofile=coverage.out ./... && \
		echo "" && \
		go tool cover -func=coverage.out
