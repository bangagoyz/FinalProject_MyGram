package repository

import (
	"finalProject/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Add(newUser model.User) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetByUsername(username string) (model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Add(newUser model.User) (model.User, error) {
	tx := ur.db.Create(&newUser)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return newUser, nil
}

func (ur *UserRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	tx := ur.db.First(&user, "email = ?", email)
	return user, tx.Error
}

func (ur *UserRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	tx := ur.db.First(&user, "username = ?", username)
	return user, tx.Error
}
