package service

import (
	"hacktiv8-msib-final-project-2/dto"
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/socialmedia_repository"
	"hacktiv8-msib-final-project-2/repository/user_repository"
)

type SocialMediaService interface {
	CreateSocialMedia(user *entity.User, payload *dto.CreateSocialMediaRequest) (*dto.CreateSocialMediaResponse, errs.MessageErr)
}

type socialmediaService struct {
	socialmediaRepo socialmedia_repository.SocialMediaRepository
	userRepo        user_repository.UserRepository
}

func NewSocialMediaService(
	socialmediaRepo socialmedia_repository.SocialMediaRepository,
	userRepo user_repository.UserRepository,
) SocialMediaService {
	return &socialmediaService{socialmediaRepo: socialmediaRepo, userRepo: userRepo}
}

func (s *socialmediaService) CreateSocialMedia(user *entity.User, payload *dto.CreateSocialMediaRequest) (*dto.CreateSocialMediaResponse, errs.MessageErr) {
	socialmedia := payload.ToEntity()

	createdSocialMedia, err := s.socialmediaRepo.CreateSocialMedia(user, socialmedia)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateSocialMediaResponse{
		ID:             createdSocialMedia.ID,
		Name:           createdSocialMedia.Name,
		SocialMediaURL: createdSocialMedia.SocialMediaURL,
		UserID:         createdSocialMedia.UserID,
		CreatedAt:      createdSocialMedia.CreatedAt,
	}

	return response, nil
}
