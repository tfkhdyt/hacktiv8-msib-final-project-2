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

type GetAllSocialMediasResponse struct {
	SocialMedias []SocialMediaData `json:"social_medias"`
}

type UserDataSos struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
}

func (s *UpdateSocialMediaRequest) ToEntity() *entity.SocialMedia {
	return &entity.SocialMedia{
		Name:           s.Name,
		SocialMediaURL: s.SocialMediaURL,
	}
}

type UpdateSocialMediaResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaData struct {
	ID             uint        `json:"id"`
	Name           string      `json:"name"`
	SocialMediaURL string      `json:"social_media_url"`
	UserID         uint        `json:"user_id"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	User           UserDataSos `json:"user"`
}

type DeleteSocialMediaResponse struct {
	Message string `json:"message"`
}
