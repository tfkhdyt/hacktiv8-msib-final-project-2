package service

import (
	"hacktiv8-msib-final-project-2/dto"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/photo_repository"
)

type PhotoService interface {
	CreatePhoto(payload *dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr)
}

type photoService struct {
	photoRepo photo_repository.PhotoRepository
}

func NewPhotoService(photoRepo photo_repository.PhotoRepository) PhotoService {
	return &photoService{photoRepo: photoRepo}
}

func (p *photoService) CreatePhoto(payload *dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr) {
	photo := payload.ToEntity()

	createdPhoto, err := p.photoRepo.CreatePhoto(payload.UserID, photo)
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
