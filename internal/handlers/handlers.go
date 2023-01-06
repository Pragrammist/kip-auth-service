package handlers

import (
	"context"
	"fmt"
	"net/http"

	proto "auth-service/pkg/proto"

	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/protobuf/types/known/timestamppb"

	"encoding/json"
	"github.com/labstack/echo/v4"
)

type HandlersBase struct {
	Client proto.ProfileClient
}

func NewHandlers(cl proto.ProfileClient) *HandlersBase {
	return &HandlersBase{Client: cl}
}

// Auth by email
// @Summary Authentification using email address
// @Tags         auth
// @Accept       json
// @Produce      json
// @Router /auth/email [post]
// @Success 200 {object} models.UserResponse
func (hb *HandlersBase) EmailAuth(c echo.Context) error {
	req := proto.CreateProfileRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&req)

	_, err = hb.Client.CreateProfile(context.Background(), &req)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c.JSON(http.StatusOK, req)
}
