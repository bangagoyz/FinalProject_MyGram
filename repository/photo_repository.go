package repository

import (
	"finalProject/model"

	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPhotoRepository interface {
	Add(newPhoto model.Photo) error
	FindAll() ([]model.Photo, error)
	GetOne(photoID string) (model.Photo, error)
	PhotoUpdate(request model.Photo, photoID string) (model.Photo, error)
	DeletePhoto(PhotoId string) error
}

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		db: db,
	}
}

func (pr *PhotoRepository) Add(newPhoto model.Photo) error {
	tx := pr.db.Create(&newPhoto)
	return tx.Error
}

func (pr *PhotoRepository) FindAll() ([]model.Photo, error) {
	photos := []model.Photo{}

	tx := pr.db.Find(&photos)
	return photos, tx.Error
}

func (pr *PhotoRepository) GetOne(photoID string) (model.Photo, error) {
	photo := model.Photo{}

	err := pr.db.Debug().Where("photo_id = ?", photoID).Take(&photo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Photo{}, model.ErrorNotFound
	}

	return photo, err
}

func (pr *PhotoRepository) PhotoUpdate(request model.Photo, photoID string) (model.Photo, error) {
	err := pr.db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{Name: "photo_id"},
			{Name: "user_id"},
			{Name: "updated_at"},
		},
	},
	).Where("photo_id = ?", photoID).Updates(&request)

	return request, err.Error
}

func (pr *PhotoRepository) DeletePhoto(PhotoId string) error {
	delPhoto := model.Photo{
		PhotoID: PhotoId,
	}

	tx := pr.db.Select("Comments").Delete(&delPhoto)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
