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
	GetAllSocialMediasByUserSosmed(userID uint) (*dto.GetAllSocialMediasResponse, errs.MessageErr)
	UpdateSocialMedia(id uint, payload *dto.UpdateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr)
	DeleteSocialMedia(id uint) (*dto.DeleteSocialMediaResponse, errs.MessageErr)
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

func (s *socialmediaService) GetAllSocialMediasByUserSosmed(userID uint) (*dto.GetAllSocialMediasResponse, errs.MessageErr) {
	socialmedias, err := s.socialmediaRepo.GetAllSocialMediasByUserSosmed(userID)
	if err != nil {
		return nil, err
	}

	socialMediasData := []dto.SocialMediaData{}
	for _, socialmedia := range socialmedias {
		user, err := s.userRepo.GetUserByID(socialmedia.UserID)
		if err != nil {
			return nil, err
		}

		socialMediasData = append(socialMediasData, dto.SocialMediaData{
			ID:             socialmedia.ID,
			Name:           socialmedia.Name,
			SocialMediaURL: socialmedia.SocialMediaURL,
			UserID:         socialmedia.UserID,
			CreatedAt:      socialmedia.CreatedAt,
			UpdatedAt:      socialmedia.UpdatedAt,
			User: dto.UserDataSos{
				ID:       user.ID,
				Username: user.Username,
			},
		},
		)
	}

	response := &dto.GetAllSocialMediasResponse{
		SocialMedias: socialMediasData,
	}

	return response, nil
}

func (s *socialmediaService) UpdateSocialMedia(id uint, payload *dto.UpdateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr) {
	oldSocialMedia, err := s.socialmediaRepo.GetSocialMediaByID(id)
	if err != nil {
		return nil, err
	}

	newSocialMedia := payload.ToEntity()

	updatedSocialMedia, err2 := s.socialmediaRepo.UpdateSocialMedia(oldSocialMedia, newSocialMedia)
	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateSocialMediaResponse{
		ID:             updatedSocialMedia.ID,
		Name:           updatedSocialMedia.Name,
		SocialMediaURL: updatedSocialMedia.SocialMediaURL,
		UserID:         updatedSocialMedia.UserID,
		UpdatedAt:      updatedSocialMedia.UpdatedAt,
	}

	return response, nil
}

func (s *socialmediaService) DeleteSocialMedia(id uint) (*dto.DeleteSocialMediaResponse, errs.MessageErr) {
	if err := s.socialmediaRepo.DeleteSocialMedia(id); err != nil {
		return nil, err
	}

	response := &dto.DeleteSocialMediaResponse{
		Message: "Your social media has been successfully deleted",
	}

	return response, nil
}
