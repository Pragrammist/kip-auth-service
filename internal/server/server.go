package server

import (
	"auth-service/internal/handlers"
	"auth-service/pkg/logger"
	kip "auth-service/pkg/proto"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/vk"
	"github.com/markbates/goth/providers/yandex"
	echoSwagger "github.com/swaggo/echo-swagger"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "auth-service/api/swagger"
)

// @title Auth service API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func Run() {
	log := logger.NewApiLogger()

	port := os.Getenv("PORT")
	mailerHost := os.Getenv("MAILER_HOST")
	profileHost := os.Getenv("PROFILE_HOST")

	// >???
	conn, err := grpc.Dial(profileHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := kip.NewProfileClient(conn)

	mailer, err := grpc.Dial(mailerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	mailerClient := kip.NewMailerServiceClient(mailer)

	if err != nil {
		fmt.Println(err.Error())
	}

	url := fmt.Sprintf("http://localhost:%s/auth", port)

	//goth.UseProviders(
	//	google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), url+"/google/callback"),
	//	yandex.New(os.Getenv("YANDEX_KEY"), os.Getenv("YANDEX_SECRET"), url+"/yandex/callback"),
	//	vk.New(os.Getenv("VK_KEY"), os.Getenv("VK_SECRET"), url+"/vk/callback"),
	//)

	goth.UseProviders(
		google.New("331459395488-mbet60ne1bod19lk484agchfpmcmdb6c.apps.googleusercontent.com", "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", url+"/google/callback"),
		yandex.New("6e452e8e71144b028fb7f30eb6ab8cc8", "0d81df3235ca4967b13ea15ff8fe2038", url+"/yandex/callback"),
		vk.New("51591323", "5AKxW73RaFMOTKHC4jcD", url+"/vk/callback"),
	)

	hand := handlers.NewHandlers(client, mailerClient)

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/auth/email", hand.EmailAuth)
	e.GET("/auth/:provider/", hand.OauthProviderAuth)
	e.GET("/auth/:provider/callback", hand.OauthProviderCallback)

	log.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
