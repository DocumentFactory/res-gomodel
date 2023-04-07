package fileshare

import (
	"context"
	"time"

	"github.com/pnocera/res-gomodel/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

func (c *AuthClient) Login() (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "fileshare")

	req := &pb.LoginRequest{
		Username: c.username,
		Password: c.password,
	}

	res, err := c.service.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil

}
