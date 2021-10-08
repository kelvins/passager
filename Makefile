build:
	@go build -o passager cmd/passager/main.go

tests:
	@go test ./...
