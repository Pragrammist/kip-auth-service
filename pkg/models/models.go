package models

import kip "auth-service/pkg/proto"

// simple swapers for swaggo/swag package

type UserRequest struct {
	kip.CreateProfileRequest
}

type UserResponse struct {
	kip.ProfileResponse
}
