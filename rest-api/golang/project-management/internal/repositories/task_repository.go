package repositories

import (
	"github.com/JawherKl/project-management/internal/db"
	"github.com/JawherKl/project-management/internal/models"
)

func CreateTask(task *models.Task) error {
	return db.DB.Create(task).Error
}

func GetTasksByProject(projectID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := db.DB.Where("project_id = ?", projectID).Find(&tasks).Error
	return tasks, err
}