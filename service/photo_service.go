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
	UpdatePhoto(id uint, payload *dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr)
	DeletePhoto(id uint) (*dto.DeletePhotoResponse, errs.MessageErr)
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

func (p *photoService) UpdatePhoto(id uint, payload *dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr) {
	oldPhoto, err := p.photoRepo.GetPhotoByID(id)
	if err != nil {
		return nil, err
	}
	newPhoto := payload.ToEntity()

	updatedPhoto, err2 := p.photoRepo.UpdatePhoto(oldPhoto, newPhoto)
	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdatePhotoResponse{
		ID:        updatedPhoto.ID,
		Title:     updatedPhoto.Title,
		Caption:   updatedPhoto.Caption,
		PhotoURL:  updatedPhoto.PhotoURL,
		UserID:    updatedPhoto.UserID,
		UpdatedAt: updatedPhoto.UpdatedAt,
	}

	return response, nil
}

func (p *photoService) DeletePhoto(id uint) (*dto.DeletePhotoResponse, errs.MessageErr) {
	if err := p.photoRepo.DeletePhoto(id); err != nil {
		return nil, err
	}

	response := &dto.DeletePhotoResponse{
		Message: "Your photo has been successfully deleted",
	}

	return response, nil
}
