package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `binding:"required" json:"first_name" validate:"required,min=3,max=20"`
	LastName  string    `binding:"required" json:"last_name" validate:"required,min=3,max=20"`
	Email     string    `binding:"required" json:"email" validate:"required,email"`
	Password  string    `binding:"required" json:"password_hash" validate:"required,min=8,strong_password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) HashPassword(plaintextPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(plaintextPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plaintextPassword))
}
