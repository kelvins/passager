build:
	@go build -o passager cmd/passager/main.go

tests:
	@go test -coverprofile=coverage.out ./... && \
		echo "" && \
		go tool cover -func=coverage.out
