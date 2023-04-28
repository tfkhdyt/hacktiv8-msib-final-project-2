package entity

import (
	"hacktiv8-msib-final-project-2/pkg/errs"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" binding:"required"`
	Email    string `gorm:"unique;not null" binding:"email,required"`
	Password string `gorm:"not null" binding:"required,min=6"`
	Age      uint   `gorm:"not null" binding:"required,min=8"`
}

func (u *User) HashPassword() errs.MessageErr {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return errs.NewInternalServerError("Failed to hash password")
	}

	u.Password = string(hashedPassword)

	return nil
}
