package handlers

import (
	"auth-service/internal/utils"
	"context"
	"fmt"
	"net/http"

	kip "auth-service/pkg/proto"
	proto "auth-service/pkg/proto"

	_ "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"
	_ "google.golang.org/protobuf/types/known/timestamppb"

	"encoding/json"

	"github.com/labstack/echo/v4"
)

type HandlersBase struct {
	Client       proto.ProfileClient
	MailerClient proto.MailerServiceClient
}

func NewHandlers(cl proto.ProfileClient, ml proto.MailerServiceClient) *HandlersBase {
	return &HandlersBase{Client: cl, MailerClient: ml}
}

func (hb *HandlersBase) OauthProviderAuth(c echo.Context) error {
	// utils for gothic
	gothic.GetProviderName = func(r *http.Request) (string, error) { return c.Param("provider"), nil }
	res := c.Response().Writer
	req := c.Request()

	if user, err := gothic.CompleteUserAuth(res, req); err == nil {
		// if user logged in before with this provider
		userReq := proto.CreateProfileRequest{
			Email:    user.Email,
			Login:    utils.GetFirstPartOfEmail(user.Email),
			Password: uuid.New().String(),
		}

		_, err = hb.Client.CreateProfile(context.Background(), &userReq)

		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		gothic.BeginAuthHandler(res, req)
	}

	return c.String(http.StatusOK, "OK")
}

func (hb *HandlersBase) OauthProviderCallback(c echo.Context) error {
	res := c.Response().Writer
	req := c.Request()

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Println(err.Error())
	}

	userReq := proto.CreateProfileRequest{
		Email:    user.Email,
		Login:    utils.GetFirstPartOfEmail(user.Email),
		Password: uuid.New().String(),
	}

	_, err = hb.Client.CreateProfile(context.Background(), &userReq)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c.String(http.StatusOK, "OK")
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

	email := &kip.EmailRequest{To: []string{"ignat.belousov2005@gmail.com"}, ContentType: "text/plain", Subject: "Ваш код подтверждения аккаунта KIP", Body: "12344123"}
	_, err = hb.MailerClient.SendEmails(context.Background(), email)

	//_, err = hb.Client.CreateProfile(context.Background(), &req)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c.JSON(http.StatusOK, req)
}
