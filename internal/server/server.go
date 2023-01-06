package server

import (
	"auth-service/internal/handlers"
	"auth-service/pkg/logger"
	kip "auth-service/pkg/proto"
	"fmt"

	"github.com/labstack/echo/v4"
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

	conn, err := grpc.Dial("localhost:81", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := kip.NewProfileClient(conn)

	if err != nil {
		fmt.Println(err.Error())
	}

	hand := handlers.NewHandlers(client)

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/auth/email", hand.EmailAuth)

	log.Fatal(e.Start(":5100"))
}
