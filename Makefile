# note: call scripts from /scripts

run:
	@go run ./cmd/auth_service/main.go

tests:
	go test -v ./test/...

clean:
	go mod tidy && go fmt ./...

lint:
	golangci-lint run ./...

swag:
	@swag init --parseInternal -o api/swagger -g ./cmd/auth-service/main.go
