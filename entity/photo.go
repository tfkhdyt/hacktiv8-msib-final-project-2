package entity

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" gorm:"not null"`
	UserID   uint   `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments []Comment
}
