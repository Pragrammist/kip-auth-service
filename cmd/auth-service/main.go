package main

import (
	"auth-service/internal/server"

	// swaggo doesn't work without that (issue #818)
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	server.Run()
}
