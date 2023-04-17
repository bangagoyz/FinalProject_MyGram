package repository

import (
	"errors"
	"finalProject/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ISocialMediaRepository interface {
	Add(newSocial model.SocialMedia) error
	Get() ([]model.SocialMedia, error)
	GetOne(SocialID string) (model.SocialMedia, error)
	Update(updateSocialMedia model.SocialMedia, socialId string) (model.SocialMedia, error)
	Delete(socialID string) error
}

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{
		db: db,
	}
}

func (sr *SocialMediaRepository) Add(newSocial model.SocialMedia) error {
	tx := sr.db.Create(&newSocial)
	return tx.Error
}

func (sr *SocialMediaRepository) Get() ([]model.SocialMedia, error) {
	socMed := []model.SocialMedia{}

	tx := sr.db.Find(&socMed)
	return socMed, tx.Error
}

func (sr *SocialMediaRepository) GetOne(SocialID string) (model.SocialMedia, error) {
	FindSocial := model.SocialMedia{}

	err := sr.db.Debug().Where("social_id = ?", SocialID).Take(&FindSocial).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.SocialMedia{}, model.ErrorNotFound
	}
	return FindSocial, err
}

func (sr *SocialMediaRepository) Update(updateSocialMedia model.SocialMedia, socialId string) (model.SocialMedia, error) {
	err := sr.db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{Name: "social_id"},
			{Name: "user_id"},
			{Name: "updated_at"},
		},
	},
	).
		Where("social_id = ?", socialId).Updates(&updateSocialMedia)
	return updateSocialMedia, err.Error
}

func (sr *SocialMediaRepository) Delete(socialID string) error {
	delSocial := model.SocialMedia{}

	tx := sr.db.Delete(&delSocial, "social_id = ?", socialID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
