package service

import (
	"context"
	"fmt"

	userpb "github.com/arslan-atajykov/ecommerce-platform/proto/userpb"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
	DB *pgxpool.Pool
}

func NewUserServiceServer(db *pgxpool.Pool) *UserServiceServer {
	return &UserServiceServer{DB: db}
}

func (s *UserServiceServer) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	var id string
	err = s.DB.QueryRow(ctx, "INSERT INTO users (email, password, name) VALUES($1,$2,$3) RETURNING id", req.Email, string(hashed), req.Name).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	user := &userpb.User{
		Id:    id,
		Email: req.Email,
		Name:  req.Name,
	}
	fmt.Println("Register called:", req.Email)
	return &userpb.RegisterResponse{User: user}, nil

}

func (s *UserServiceServer) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	var id, hashed, name string
	err := s.DB.QueryRow(ctx, "SELECT id, password, name FROM users WHERE email=$1", req.Email).Scan(&id, &hashed, &name)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(req.Password)) != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// JWT generation will come later
	token := "dummy-jwt-token"

	return &userpb.LoginResponse{
		Token: token,
		User: &userpb.User{
			Id:    id,
			Email: req.Email,
			Name:  name,
		},
	}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	var email, name string
	err := s.DB.QueryRow(ctx, "SELECT email, name FROM users WHERE id=$1", req.Id).Scan(&email, &name)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    req.Id,
			Email: email,
			Name:  name,
		},
	}, nil
}
