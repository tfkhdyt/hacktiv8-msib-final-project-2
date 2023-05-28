package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

var ginMode = os.Getenv("GIN_MODE")

func init() {
	if ginMode != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func GetDBConfig() gorm.Dialector {
	dbConfig := databaseConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbName:   os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConfig.host,
		dbConfig.port,
		dbConfig.user,
		dbConfig.password,
		dbConfig.dbName,
	)

	return postgres.Open(dsn)
}
