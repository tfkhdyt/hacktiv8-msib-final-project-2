package database

import (
	"log"

	"gorm.io/gorm"

	"hacktiv8-msib-final-project-2/config"
	"hacktiv8-msib-final-project-2/entity"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(config.GetDBConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err = db.AutoMigrate(&entity.User{}, &entity.Photo{}, &entity.Comment{}, &entity.SocialMedia{}); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Connected to DB!")
}

func GetPostgresInstance() *gorm.DB {
	return db
}
