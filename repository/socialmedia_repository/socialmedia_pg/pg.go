package socialmedia_pg

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/socialmedia_repository"
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

func (s *socialmediaPG) GetAllSocialMediasByUserSosmed(userID uint) ([]entity.SocialMedia, errs.MessageErr) {
	var socialmedias []entity.SocialMedia
	if err := s.db.Find(&socialmedias, "user_id = ?", userID).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to get all social media")
	}

	return socialmedias, nil
}

func (s *socialmediaPG) GetSocialMediaByID(id uint) (*entity.SocialMedia, errs.MessageErr) {
	var socialmedia entity.SocialMedia
	if err := s.db.First(&socialmedia, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Social Media with id %d is not found", id))
	}

	return &socialmedia, nil
}

func (s *socialmediaPG) UpdateSocialMedia(oldSocialMedia *entity.SocialMedia, newSocialMedia *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	if err := s.db.Model(oldSocialMedia).Updates(newSocialMedia).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update social media with id %d", oldSocialMedia.ID))
	}

	return oldSocialMedia, nil
}

func (s *socialmediaPG) DeleteSocialMedia(id uint) errs.MessageErr {
	if err := s.db.Delete(&entity.SocialMedia{}, id).Error; err != nil {
		log.Println("Error:", err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete Social Media with id %d", id))
	}

	return nil
}
