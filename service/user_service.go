package service

import (
	"hacktiv8-msib-final-project-2/dto"
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/user_repository"
)

type UserService interface {
	Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	UpdateUser(userID uint, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr)
	DeleteUser(id uint) (*dto.DeleteUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {
	user := payload.ToEntity()
	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	registeredUser, err := u.userRepo.Register(user)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		ID:       registeredUser.ID,
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
		Age:      registeredUser.Age,
	}

	return response, nil
}

func (u *userService) Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := user.ComparePassword(payload.Password); err != nil {
		return nil, err
	}

	token, err2 := user.CreateToken()
	if err2 != nil {
		return nil, err2
	}

	response := &dto.LoginResponse{
		Token: token,
	}

	return response, nil
}

func (u *userService) UpdateUser(userID uint, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr) {
	oldUser, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	newUser := &entity.User{
		Email:    payload.Email,
		Username: payload.Username,
	}

	updatedUser, err2 := u.userRepo.UpdateUser(oldUser, newUser)
	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateUserResponse{
		ID:        updatedUser.ID,
		Email:     updatedUser.Email,
		Username:  updatedUser.Username,
		Age:       updatedUser.Age,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteUser(id uint) (*dto.DeleteUserResponse, errs.MessageErr) {
	if err := u.userRepo.DeleteUser(id); err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{
		Message: "Your account has been successfully deleted",
	}

	return response, nil
}
