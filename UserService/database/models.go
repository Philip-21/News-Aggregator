package database

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"-" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//The Models struct provides access to the User struct, which can be used throughout the application codebase.
type Models struct{
	Users User
}

