package test

import (
	pb "auth-service/pkg/proto"
	"fmt"

	"context"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (pb.ProfileClient, func()) {
	buffer := 1024 * 1024
	listener := bufconn.Listen(buffer)

	s := grpc.NewServer()
	go func() {
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.Dial("localhost:81", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err.Error())
	}

	client := pb.NewProfileClient(conn)

	return client, s.Stop
}

func TestGrpcServer_CreateProfile(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)

	defer closer()

	f := uuid.New()

	out, err := client.CreateProfile(ctx, &pb.CreateProfileRequest{
		Email:    fmt.Sprintf("%s@gmail.com", f),
		Login:    fmt.Sprintf("politech%s", f),
		Password: "123124124124",
	})

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.True(t, out.User.Login == fmt.Sprintf("politech%s", f))
}
