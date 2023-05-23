package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `gorm:"not null" json:"message"`
}
