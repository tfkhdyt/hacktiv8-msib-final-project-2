package service

import (
	"hacktiv8-msib-final-project-2/dto"
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/photo_repository"
	"hacktiv8-msib-final-project-2/repository/user_repository"
)

type PhotoService interface {
	CreatePhoto(user *entity.User, payload *dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr)
	GetAllPhotos() ([]dto.GetAllPhotosResponse, errs.MessageErr)
}

type photoService struct {
	photoRepo photo_repository.PhotoRepository
	userRepo  user_repository.UserRepository
}

func NewPhotoService(photoRepo photo_repository.PhotoRepository, userRepo user_repository.UserRepository) PhotoService {
	return &photoService{photoRepo: photoRepo, userRepo: userRepo}
}

func (p *photoService) CreatePhoto(user *entity.User, payload *dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr) {
	photo := payload.ToEntity()

	createdPhoto, err := p.photoRepo.CreatePhoto(user, photo)
	if err != nil {
		return nil, err
	}

	response := &dto.CreatePhotoResponse{
		ID:        createdPhoto.ID,
		Title:     createdPhoto.Title,
		Caption:   createdPhoto.Caption,
		PhotoURL:  createdPhoto.PhotoURL,
		UserID:    createdPhoto.UserID,
		CreatedAt: createdPhoto.CreatedAt,
	}

	return response, nil
}

func (p *photoService) GetAllPhotos() ([]dto.GetAllPhotosResponse, errs.MessageErr) {
	photos, err := p.photoRepo.GetAllPhotos()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllPhotosResponse{}
	for _, photo := range photos {
		user, err := p.userRepo.GetUserByID(photo.UserID)
		if err != nil {
			return nil, err
		}

		response = append(response, dto.GetAllPhotosResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: dto.UserData{
				Email:    user.Email,
				Username: user.Username,
			},
		})
	}

	return response, nil
}
