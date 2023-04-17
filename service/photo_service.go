package service

import (
	"finalProject/helper"
	"finalProject/model"
	"finalProject/repository"
)

type IPhotoService interface {
	Create(request model.PhotoRequest, userID string) (model.PhotoCreateResponse, error)
	GetAllPhoto() ([]model.PhotoResponse, error)
	GetOnePhoto(photoID string) (model.PhotoResponse, error)
	DeletePhoto(photoID string, userID string) error
	UpdatePhoto(request model.PhotoRequest, userID string, photoID string) (model.PhotoResponse, error)
}

type PhotoService struct {
	PhotoRepository   repository.IPhotoRepository
	CommentRepository repository.ICommentRepository
}

func NewPhotoService(photoRepository repository.IPhotoRepository, Commentrepository repository.ICommentRepository) *PhotoService {
	return &PhotoService{
		PhotoRepository:   photoRepository,
		CommentRepository: Commentrepository,
	}
}

func (ps *PhotoService) Create(request model.PhotoRequest, userID string) (model.PhotoCreateResponse, error) {
	PhotoID := helper.GenerateID()

	NewPhoto := model.Photo{
		PhotoID:  PhotoID,
		Title:    request.Title,
		PhotoUrl: request.PhotoUrl,
		UserID:   userID,
	}

	err := ps.PhotoRepository.Add(NewPhoto)
	if err != nil {
		if err != model.ErrorNotFound {
			return model.PhotoCreateResponse{}, err
		}

		return model.PhotoCreateResponse{}, model.ErrorNotFound
	}

	response := model.PhotoCreateResponse{
		PhotoID:   NewPhoto.PhotoID,
		Title:     NewPhoto.Title,
		PhotoUrl:  NewPhoto.PhotoUrl,
		UserID:    NewPhoto.UserID,
		CreatedAt: NewPhoto.CreatedAt,
	}
	return response, nil
}

func (ps *PhotoService) GetAllPhoto() ([]model.PhotoAllResponse, error) {
	photoResults := []model.PhotoAllResponse{}

	res, err := ps.PhotoRepository.FindAll()

	if err != nil {
		return []model.PhotoAllResponse{}, err
	}

	for _, reqRes := range res {
		photoResults = append(photoResults, model.PhotoAllResponse{
			PhotoID:   reqRes.PhotoID,
			Title:     reqRes.Title,
			PhotoUrl:  reqRes.PhotoUrl,
			UserID:    reqRes.UserID,
			CreatedAt: reqRes.CreatedAt,
			UpdatedAt: reqRes.UpdatedAt,
		})
	}

	return photoResults, nil
}

func (ps *PhotoService) GetOnePhoto(photoID string) (model.PhotoResponse, error) {
	photoRequest, err := ps.PhotoRepository.GetOne(photoID)
	if err != nil {
		if err == model.ErrorNotFound {
			return model.PhotoResponse{}, err
		}

		return model.PhotoResponse{}, model.ErrorNotFound
	}

	comments := []model.Comment{}
	commentsResp, err := ps.CommentRepository.FindCommentByPhoto(photoID)
	for _, comment := range commentsResp {
		comments = append(comments, model.Comment(comment))
	}
	if err != nil {
		return model.PhotoResponse{}, err
	}
	return model.PhotoResponse{
		PhotoID:   photoRequest.PhotoID,
		Title:     photoRequest.Title,
		PhotoUrl:  photoRequest.PhotoUrl,
		UserID:    photoRequest.UserID,
		Comments:  comments,
		CreatedAt: photoRequest.CreatedAt,
		UpdatedAt: photoRequest.UpdatedAt,
	}, nil
}

func (ps *PhotoService) UpdatePhoto(request model.PhotoRequest, userID string, photoID string) (model.PhotoResponse, error) {
	findPhoto, err := ps.PhotoRepository.GetOne(photoID)
	if err != nil {
		if err != model.ErrorNotFound {
			return model.PhotoResponse{}, err
		}
		return model.PhotoResponse{}, model.ErrorNotFound
	}

	if userID != findPhoto.UserID {
		return model.PhotoResponse{}, model.ErrorForbiddenAccess
	}

	updateReq := model.Photo{
		Title:    request.Title,
		PhotoUrl: request.PhotoUrl,
	}

	res, err := ps.PhotoRepository.PhotoUpdate(updateReq, photoID)
	if err != nil {
		return model.PhotoResponse{}, err
	}

	return model.PhotoResponse{
		PhotoID:   res.PhotoID,
		Title:     res.Title,
		PhotoUrl:  res.PhotoUrl,
		UserID:    res.UserID,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (ps *PhotoService) DeletePhoto(photoID string, userID string) error {
	findPhoto, err := ps.PhotoRepository.GetOne(photoID)
	if err != nil {
		if err != model.ErrorNotFound {
			return err
		}
		return model.ErrorNotFound
	}

	if userID != findPhoto.UserID {
		return model.ErrorForbiddenAccess
	}

	err = ps.PhotoRepository.DeletePhoto(photoID)
	if err != nil {
		return err
	}

	return nil
}
