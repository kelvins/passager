build:
	@go build -o passager cmd/passager/main.go

format:
	@go fmt ./...

tests:
	@go test -coverprofile=coverage.out ./... && \
		echo "" && \
		go tool cover -func=coverage.out
