# note: call scripts from /scripts

run:
	@go run ./cmd/auth-service/main.go

swag:
	@swag init -o api/swagger -g ./cmd/auth-service/main.go
