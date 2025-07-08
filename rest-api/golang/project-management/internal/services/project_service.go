package services

import (
	"github.com/JawherKl/project-management/internal/models"
	"github.com/JawherKl/project-management/internal/repositories"
)

func GetProjects() ([]models.Project, error) {
	return repositories.GetAllProjects()
}

func CreateProject(p *models.Project) error {
	return repositories.CreateProject(p)
}
