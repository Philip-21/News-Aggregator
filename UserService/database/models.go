package database

import (
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID        int      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"-" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// The Models struct provides access to the User struct, which can be used throughout the application codebase.
type Models struct {
	Users User
}

// used to create an instance of the database package by connecting to the db.
// It returns the type Model, which embeds all the types we want to be available to our application.
func New(dbPool *gorm.DB) Models {
	db = dbPool

	return Models{
		Users: User{},
	}
}
