package service

import (
	"finalProject/helper"
	"finalProject/model"
	"finalProject/repository"
)

type iCommentService interface {
	CreateComment(request model.CommentCreateRequest, userID string, photoID string) (*model.CommentCreateResponse, error)
	GetAll() ([]model.CommentResponse, error)
	Update(UpdateComment model.CommentUpdateRequest, CommentID string, userID string) (model.CommentUpdateResponse, error)
	GetOne(commentID string) (model.CommentResponse, error)
	Delete(commentID string, userID string) error
}

type CommentService struct {
	CommentRepository repository.ICommentRepository
	PhotoRepository   repository.IPhotoRepository
}

func NewCommentService(commentRepository repository.ICommentRepository, PhotoReposit repository.IPhotoRepository) *CommentService {
	return &CommentService{
		CommentRepository: commentRepository,
		PhotoRepository:   PhotoReposit,
	}
}

func (cs *CommentService) CreateComment(request model.CommentCreateRequest, userID string, photoID string) (*model.CommentCreateResponse, error) {
	_, err := cs.PhotoRepository.GetOne(photoID)
	if err != nil {
		if err != model.ErrorNotFound {
			return &model.CommentCreateResponse{}, err
		}
		return &model.CommentCreateResponse{}, model.ErrorNotFound
	}

	commentID := helper.GenerateID()

	NewComment := model.Comment{
		CommentID: commentID,
		Message:   request.Message,
		UserID:    userID,
		PhotoID:   photoID,
	}

	err = cs.CommentRepository.CreateComment(NewComment)
	if err != nil {
		if err != model.ErrorNotFound {
			return &model.CommentCreateResponse{}, err
		}
		return &model.CommentCreateResponse{}, model.ErrorNotFound
	}

	return &model.CommentCreateResponse{
		CommentID: NewComment.CommentID,
		Message:   NewComment.Message,
		PhotoID:   NewComment.PhotoID,
		UserID:    NewComment.UserID,
		CreatedAt: NewComment.CreatedAt,
	}, err

}

func (cs *CommentService) GetAll() ([]model.CommentResponse, error) {
	AllComment := []model.CommentResponse{}

	res, err := cs.CommentRepository.Get()
	if err != nil {
		return []model.CommentResponse{}, err
	}
	for _, CommentRes := range res {
		AllComment = append(AllComment, model.CommentResponse{
			CommentID: CommentRes.CommentID,
			Message:   CommentRes.Message,
			UserID:    CommentRes.UserID,
			PhotoID:   CommentRes.PhotoID,
			CreatedAt: CommentRes.CreatedAt,
			UpdatedAt: CommentRes.UpdatedAt,
		})
	}

	return AllComment, nil

}
func (cs *CommentService) Update(UpdateComment model.CommentUpdateRequest, CommentID string, userID string) (model.CommentUpdateResponse, error) {
	getID, err := cs.CommentRepository.GetOne(CommentID)
	if err != nil {
		if err != model.ErrorNotFound {
			return model.CommentUpdateResponse{}, err
		}
		return model.CommentUpdateResponse{}, model.ErrorNotFound
	}

	if getID.UserID != userID {
		return model.CommentUpdateResponse{}, model.ErrorForbiddenAccess
	}

	CommentUpdate := model.Comment{
		Message: UpdateComment.Message,
	}

	res, err := cs.CommentRepository.Update(CommentUpdate, CommentID)

	if err != nil {
		return model.CommentUpdateResponse{}, err

	}

	return model.CommentUpdateResponse{
		CommentID: res.CommentID,
		Message:   res.Message,
		UserID:    res.UserID,
		PhotoID:   res.PhotoID,
		UpdatedAt: res.UpdatedAt,
	}, nil

}

func (cs *CommentService) GetOne(commentID string) (model.CommentResponse, error) {
	getOne, err := cs.CommentRepository.GetOne(commentID)

	if err != nil {
		if err != model.ErrorNotFound {
			return model.CommentResponse{}, err

		}
		return model.CommentResponse{}, model.ErrorNotFound
	}

	return model.CommentResponse{
		CommentID: getOne.CommentID,
		Message:   getOne.Message,
		UserID:    getOne.UserID,
		PhotoID:   getOne.PhotoID,
		CreatedAt: getOne.CreatedAt,
		UpdatedAt: getOne.UpdatedAt,
	}, nil
}

func (cs *CommentService) Delete(commentID string, userID string) error {
	getCommentID, err := cs.CommentRepository.GetOne(commentID)

	if err != nil {
		if err != model.ErrorNotFound {
			return err

		}
		return model.ErrorNotFound
	}

	if getCommentID.UserID != userID {
		return model.ErrorForbiddenAccess
	}

	err = cs.CommentRepository.Delete(commentID)

	if err != nil {
		return err
	}
	return nil
}
