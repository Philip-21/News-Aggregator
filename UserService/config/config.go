package config

import (
	"user/database"

	"github.com/alexedwards/scs/v2"
	"gorm.io/gorm"
)

type AppConfig struct {
	Session *scs.SessionManager
	DB      *gorm.DB
	Models  database.Models
}
