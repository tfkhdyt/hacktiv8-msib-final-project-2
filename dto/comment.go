package dto

import (
	"hacktiv8-msib-final-project-2/entity"
	"time"
)

type CreateCommentRequest struct {
	Message string `json:"message" binding:"required"`
	PhotoID uint   `json:"photo_id"`
}

func (c *CreateCommentRequest) ToEntity() *entity.Comment {
	return &entity.Comment{
		Message: c.Message,
		PhotoID: c.PhotoID,
	}
}

type CreateCommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllCommentsResponse struct {
	ID        uint           `json:"id"`
	Message   string         `json:"message"`
	PhotoID   uint           `json:"photo_id"`
	UserID    uint           `json:"user_id"`
	UpdateAt  time.Time      `json:"update_at"`
	CreatedAt time.Time      `json:"created_at"`
	User      UserDataWithID `json:"user"`
	Photo     PhotoData      `json:"photo"`
}

type UserDataWithID struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoData struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type UpdateCommentRequest CreateCommentRequest

func (c *UpdateCommentRequest) ToEntity() *entity.Comment {
	return &entity.Comment{
		Message: c.Message,
	}
}

type UpdateCommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}
