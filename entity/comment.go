package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint    `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PhotoID uint    `json:"photo_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Message string  `gorm:"not null" json:"message"`
}