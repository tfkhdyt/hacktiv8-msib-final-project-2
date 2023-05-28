package dto

import (
	"time"

	"hacktiv8-msib-final-project-2/entity"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required,min=6"`
	Age      uint   `json:"age" binding:"required,min=8"`
}

func (r *RegisterRequest) ToEntity() *entity.User {
	return &entity.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Age:      r.Age,
	}
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"email,required"`
	Age      uint   `json:"age" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token" binding:"jwt"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" binding:"email,required"`
	Username string `json:"username" binding:"required"`
}

type UpdateUserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email" binding:"email,required"`
	Username  string    `json:"username" binding:"required"`
	Age       uint      `json:"age" binding:"required,min=8"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserRequest = LoginResponse

type DeleteUserResponse struct {
	Message string `json:"message"`
}
