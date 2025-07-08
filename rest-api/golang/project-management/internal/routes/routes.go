package routes

import (
	"github.com/JawherKl/project-management/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/projects", handlers.GetProjects)
	api.Post("/projects", handlers.CreateProject)

	api.Post("/projects/:id/tasks", handlers.CreateTask)
	api.Get("/projects/:id/tasks", handlers.GetTasksByProject)
}
