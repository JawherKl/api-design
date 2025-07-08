package services

import (
	"github.com/JawherKl/project-management/internal/models"
	"github.com/JawherKl/project-management/internal/repositories"
)

func CreateTask(task *models.Task) error {
	return repositories.CreateTask(task)
}

func GetTasksByProject(projectID uint) ([]models.Task, error) {
	return repositories.GetTasksByProject(projectID)
}