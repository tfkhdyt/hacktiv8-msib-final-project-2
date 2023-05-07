package photo_pg

import (
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/photo_repository"
	"log"

	"gorm.io/gorm"
)

type photoPg struct {
	db *gorm.DB
}

func NewPhotoPG(db *gorm.DB) photo_repository.PhotoRepository {
	return &photoPg{db: db}
}

func (p *photoPg) CreatePhoto(user *entity.User, photo *entity.Photo) (*entity.Photo, errs.MessageErr) {
	if err := p.db.Model(user).Association("Photos").Append(photo); err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to create new photo")
	}

	return photo, nil
}

func (p *photoPg) GetAllPhotos() ([]entity.Photo, errs.MessageErr) {
	var photos []entity.Photo
	if err := p.db.Find(&photos).Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to get all photos")
	}

	return photos, nil
}
