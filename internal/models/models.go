package models

import (
	kip "auth-service/pkg/proto"

	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// simple swapers for swaggo/swag package

type UserRequest struct {
	kip.CreateProfileRequest
}

type UserResponse struct {
	kip.ProfileResponse
}
