package socialmedia_pg

import (
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/socialmedia_repository"
	"log"

	"gorm.io/gorm"
)

type socialmediaPG struct {
	db *gorm.DB
}

func NewSocialMediaPG(db *gorm.DB) socialmedia_repository.SocialMediaRepository {
	return &socialmediaPG{db: db}
}

func (s *socialmediaPG) CreateSocialMedia(user *entity.User, socialmedia *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	if err := s.db.Model(user).Association("SocialMedias").Append(socialmedia); err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to create new Social Media")
	}
	return socialmedia, nil
}
