package user_pg

import (
	"fmt"
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/user_repository"
	"log"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{db: db}
}

func (u *userPG) Register(user *entity.User) (*entity.User, errs.MessageErr) {
	if err := u.db.Create(user).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to register new user")
	}

	return user, nil
}

func (u *userPG) GetUserByEmail(email string) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with email %s is not found", email))
	}

	return &user, nil
}
