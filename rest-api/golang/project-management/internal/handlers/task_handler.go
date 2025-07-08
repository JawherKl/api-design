package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/JawherKl/project-management/internal/models"
	"github.com/JawherKl/project-management/internal/services"
	"github.com/JawherKl/project-management/internal/validators"
)

// @Summary      Get all tasks by project
// @Description  Retrieves a list of all tasks for a specific project
// @Tags         tasks
// @Produce      json
// @Success      200  {array}  models.Task
// @Router       /projects/{id}/tasks [get]
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

// @Summary      Create a new task
// @Description  Create a task with name and description
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body  models.Task  true  "Task JSON"
// @Success      201  {object}  models.Task
// @Router       /projects/{id}/tasks [post]
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

	if errs := validators.ValidateStruct(&task); errs != nil {
		return c.Status(400).JSON(fiber.Map{"validation_errors": errs})
	}

	if err := services.CreateTask(&task); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create task"})
	}
	return c.Status(201).JSON(task)
}
