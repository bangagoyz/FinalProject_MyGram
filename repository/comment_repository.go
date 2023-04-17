package repository

import (
	"errors"
	"finalProject/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ICommentRepository interface {
	CreateComment(newComment model.Comment) error
	FindCommentByPhoto(photoID string) ([]model.Comment, error)
	Get() ([]model.Comment, error)
	GetOne(CommentID string) (model.Comment, error)
	Update(UpdateComment model.Comment, CommentID string) (model.Comment, error)
	Delete(commentID string) error
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (cr *CommentRepository) CreateComment(newComment model.Comment) error {
	tx := cr.db.Create(&newComment)
	return tx.Error
}

func (cr *CommentRepository) FindCommentByPhoto(photoID string) ([]model.Comment, error) {
	comments := []model.Comment{}

	err := cr.db.Where("photo_id = ?", photoID).Find(&comments).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Comment{}, model.ErrorNotFound

		}

		return comments, err
	}

	fmt.Println("comments: ", comments)

	return comments, nil
}

func (cr *CommentRepository) Get() ([]model.Comment, error) {
	GetComment := []model.Comment{}

	tx := cr.db.Find(&GetComment)
	return GetComment, tx.Error
}

func (cr *CommentRepository) GetOne(CommentID string) (model.Comment, error) {
	getComment := model.Comment{}

	err := cr.db.Debug().Where("comment_id = ?", CommentID).Take(&getComment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Comment{}, model.ErrorNotFound
	}
	return getComment, err

}

func (cr *CommentRepository) Update(UpdateComment model.Comment, CommentID string) (model.Comment, error) {
	err := cr.db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{Name: "comment_id"},
			{Name: "user_id"},
			{Name: "updated_at"},
		},
	}).Where("comment_id = ?", CommentID).Updates(&UpdateComment)
	return UpdateComment, err.Error
}

func (cr *CommentRepository) Delete(commentID string) error {
	deleteComment := model.Comment{}

	err := cr.db.Delete(&deleteComment, "comment_id = ? ", commentID)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
