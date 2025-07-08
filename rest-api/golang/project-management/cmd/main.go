package main

import (
	"log"

    "github.com/JawherKl/project-management/config"
    "github.com/JawherKl/project-management/internal/db"
    "github.com/JawherKl/project-management/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	db.Connect(cfg.DatabaseURL)

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":" + cfg.Port))
}
