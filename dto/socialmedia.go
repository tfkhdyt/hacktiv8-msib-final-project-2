package dto

import (
	"hacktiv8-msib-final-project-2/entity"
	"time"
)

type CreateSocialMediaRequest struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
}

func (s *CreateSocialMediaRequest) ToEntity() *entity.SocialMedia {
	return &entity.SocialMedia{
		Name:           s.Name,
		SocialMediaURL: s.SocialMediaURL,
	}
}

type CreateSocialMediaResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}
