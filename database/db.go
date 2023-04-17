package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"finalProject/model"
)

var (
	host     = os.Getenv("localhost")
	user     = os.Getenv("postgres")
	password = os.Getenv("naha22")
	dbPort   = os.Getenv("5432")
	dbname   = os.Getenv("db-mygram")
	db       *gorm.DB
	err      error
)

func StartDB() {
	var err error

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
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
