package handlers

import (
	"github.com/JawherKl/project-management/internal/models"
	"github.com/JawherKl/project-management/internal/services"
	"github.com/JawherKl/project-management/internal/validators"

	"github.com/gofiber/fiber/v2"
)

func GetProjects(c *fiber.Ctx) error {
	projects, err := services.GetProjects()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch projects"})
	}
	return c.JSON(projects)
}

func CreateProject(c *fiber.Ctx) error {
	var project models.Project
	if err := c.BodyParser(&project); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if errs := validators.ValidateStruct(&project); errs != nil {
		return c.Status(400).JSON(fiber.Map{"validation_errors": errs})
	}

	if err := services.CreateProject(&project); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create project"})
	}
	return c.Status(201).JSON(project)
}
