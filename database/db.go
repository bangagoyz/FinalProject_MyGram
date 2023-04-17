package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"finalProject/model"
)

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_NAME     = os.Getenv("DB_NAME")
	db          *gorm.DB
	err         error
)

func StartDB() {
	var err error

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)
	dsn := config

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(model.User{}, model.SocialMedia{}, model.Photo{}, model.Comment{})

}
func GetDB() *gorm.DB {
	return db
}
