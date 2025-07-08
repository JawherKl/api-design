package db

import (
	"log"
	"github.com/JawherKl/project-management/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(databaseURL string) {
	var err error
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&models.Project{}, &models.Task{})
	if err != nil {
		log.Fatal("❌ Auto migration failed:", err)
	}
}
