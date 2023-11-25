package config

import (
	"user/database"

	"gorm.io/gorm"
)

type AppConfig struct {
	DB *gorm.DB
	Models database.Models
}