package repositories

import (
	"github.com/JawherKl/project-management/internal/db"
	"github.com/JawherKl/project-management/internal/models"
)

func GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	result := db.DB.Find(&projects)
	return projects, result.Error
}

func CreateProject(p *models.Project) error {
	result := db.DB.Create(p)
	return result.Error
}
