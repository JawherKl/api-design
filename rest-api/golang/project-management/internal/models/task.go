package models

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" validate:"required"`
	Status    string    `json:"status" validate:"required,oneof=todo in_progress done"`
	ProjectID uint      `json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
