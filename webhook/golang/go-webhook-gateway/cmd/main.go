package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/JawherKl/go-webhook-gateway/internal/router"
	"github.com/JawherKl/go-webhook-gateway/internal/storage"
)

func main() {
	_ = godotenv.Load()

	storage.InitMongo() // ✅ init DB

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	router.SetupRoutes(r)

	log.Printf("🚀 Webhook Gateway running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed:", err)
	}
}
