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
