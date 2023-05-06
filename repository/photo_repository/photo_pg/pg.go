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

func (p *photoPg) CreatePhoto(userID uint, photo *entity.Photo) (*entity.Photo, errs.MessageErr) {
	if err := p.db.Model(&entity.User{}).Where("id = ?", userID).Association("Photos").Append(photo); err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to create new photo")
	}

	return photo, nil
}
