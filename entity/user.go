package entity

import (
	"hacktiv8-msib-final-project-2/pkg/errs"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

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

func (u *User) ComparePassword(password string) errs.MessageErr {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return errs.NewBadRequest("Password is not valid!")
	}

	return nil
}

func (u *User) CreateToken() (string, errs.MessageErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": u.ID,
			"exp":    time.Now().Add(1 * time.Hour),
		})

	signedString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		log.Println("Error:", err.Error())
		return "", errs.NewInternalServerError("Failed to sign jwt token")
	}

	return signedString, nil
}
