package dto

import (
	"hacktiv8-msib-final-project-2/entity"
	"time"
)

type CreatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required,url"`
}

func (p *CreatePhotoRequest) ToEntity() *entity.Photo {
	return &entity.Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
	}
}

type CreatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllPhotosResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserData  `json:"user"`
}

type UserData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdatePhotoRequest CreatePhotoRequest

func (p *UpdatePhotoRequest) ToEntity() *entity.Photo {
	return &entity.Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
	}
}

type UpdatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
