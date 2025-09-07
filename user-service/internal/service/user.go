package service

import (
	"context"
	"fmt"

	userpb "github.com/arslan-atajykov/ecommerce-platform/proto/userpb"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{}
}

func (s *UserServiceServer) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {

	user := &userpb.User{
		Id:       "123",
		Email:    req.Email,
		Password: "hashed-password",
		Name:     req.Name,
	}
	fmt.Println("Register called:", req.Email)
	return &userpb.RegisterResponse{User: user}, nil

}

func (s *UserServiceServer) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	user := &userpb.User{
		Id:    "123",
		Email: req.Email,
		Name:  "Test User",
	}
	token := "dummy-jwt-token"
	fmt.Println("Login called:", req.Email)
	return &userpb.LoginResponse{Token: token, User: user}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user := &userpb.User{
		Id:    req.Id,
		Email: "dummy@example.com",
		Name:  "Dummy User",
	}
	fmt.Println("GetUser called", req.Id)
	return &userpb.GetUserResponse{User: user}, nil
}
