package database

import (
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (U *User) Insert(user User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		log.Println("Unable to hash password:", err)
		return err
	}
	newUser := &User{
		Email:     user.Email,
		Name:      user.Name,
		Password:  string(hashPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = db.Create(newUser).Error
	if err != nil {
		log.Println("Unable to create user", err)
		return err
	}
	log.Println("User created")
	// Return the new user's ID
	return nil
}

func (u *User) GetEmail(email string) (*User, error) {
	var user User
	err := db.Where("email=?", email).First(&user).Error
	if err != nil {
		log.Println("Error fetching user by email:", err)
		return nil, err
	}
	log.Println("User found for email:", email)
	return &user, nil
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			log.Println("Invalid Password")
			return false, nil
		default:
			return false, err
		}
	}
	log.Println("password validated")
	return true, nil
}
