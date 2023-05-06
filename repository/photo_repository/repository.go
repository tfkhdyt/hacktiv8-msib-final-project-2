package photo_repository

import (
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
)

type PhotoRepository interface {
	CreatePhoto(photo *entity.Photo) (*entity.Photo, errs.MessageErr)
	// GetAllPhotos() ([]entity.Photo, errs.MessageErr)
	// UpdatePhoto(oldPhoto *entity.Photo, newPhoto *entity.Photo) (*entity.Photo, errs.MessageErr)
	// DeletePhoto(id uint) errs.MessageErr
}
