package socialmedia_repository

import (
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
)

type SocialMediaRepository interface {
	CreateSocialMedia(user *entity.User, socialmedia *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr)
	GetAllSocialMediasByUserSosmed(userID uint) ([]entity.SocialMedia, errs.MessageErr)
	GetSocialMediaByID(id uint) (*entity.SocialMedia, errs.MessageErr)
	UpdateSocialMedia(oldSocialMedia *entity.SocialMedia, newSocialMedia *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr)
	DeleteSocialMedia(id uint) errs.MessageErr
}
