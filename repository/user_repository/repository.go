package user_repository

import (
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
)

type UserRepository interface {
	Register(user *entity.User) (*entity.User, errs.MessageErr)
	// Login(email, password string) (string, errs.MessageErr)
	// UpdateUser(email, password string) (*entity.User, errs.MessageErr)
	// DeleteUser(token string) errs.MessageErr
}
