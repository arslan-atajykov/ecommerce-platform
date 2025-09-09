package client

import (
	"log"

	userpb "github.com/arslan-atajykov/ecommerce-platform/proto/userpb"
	"google.golang.org/grpc"
)

type UserClient struct {
	Conn   *grpc.ClientConn
	Client userpb.UserServiceClient
}

func NewUserClient(addr string) *UserClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user-service: %v", err)
	}
	return &UserClient{
		Conn:   conn,
		Client: userpb.NewUserServiceClient(conn),
	}
}
