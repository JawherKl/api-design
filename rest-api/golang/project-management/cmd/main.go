package main

import (
	"log"

	_ "github.com/JawherKl/project-management/docs"
    "github.com/JawherKl/project-management/config"
    "github.com/JawherKl/project-management/internal/db"
    "github.com/JawherKl/project-management/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title           Project Management API
// @version         1.0
// @description     This is a simple API to manage projects and tasks.
// @host            localhost:3000
// @BasePath        /api
func main() {
	cfg := config.LoadConfig()
	db.Connect(cfg.DatabaseURL)

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":" + cfg.Port))
}
