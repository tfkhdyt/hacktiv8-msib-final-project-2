package entity

import "gorm.io/gorm"

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null"`
	SocialMediaURL string `json:"social_media_url" gorm:"not null"`
	UserID         uint   `json:"user_id"`
}
