package main

import (
	"log"
	"net"

	userpb "github.com/arslan-atajykov/ecommerce-platform/proto/userpb"
	"github.com/arslan-atajykov/ecommerce-platform/user-service/internal/service"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, service.NewUserServiceServer())

	log.Println("🚀 user-service running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
