package main

import (
	"log"

	"github.com/arslan-atajykov/ecommerce-platform/api-gateway/internal/client"
	"github.com/arslan-atajykov/ecommerce-platform/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userClient := client.NewUserClient("localhost:50051") // gRPC user-service
	userHandler := handlers.NewUserHandler(userClient)

	api := r.Group("/users")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
	}

	log.Println("🚀 api-gateway running on port 8080")
	r.Run(":8080")
}
