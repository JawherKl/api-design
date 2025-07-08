package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/JawherKl/project-management/internal/models"
	"github.com/JawherKl/project-management/internal/services"
)

func CreateTask(c *fiber.Ctx) error {
	projectID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid task input"})
	}

	task.ProjectID = uint(projectID)

	if err := services.CreateTask(&task); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create task"})
	}
	return c.Status(201).JSON(task)
}

func GetTasksByProject(c *fiber.Ctx) error {
	projectID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	tasks, err := services.GetTasksByProject(uint(projectID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch tasks"})
	}
	return c.JSON(tasks)
}