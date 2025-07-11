package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/JawherKl/go-webhook-gateway/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing...")
	}

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
