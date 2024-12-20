package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Username   string `json:"username" gorm:"not null,unique"`
	Password   string `json:"-"`
	Email      string `json:"email"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHashed)
	return nil
}
