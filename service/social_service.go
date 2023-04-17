package service

import (
	"finalProject/helper"
	"finalProject/model"
	"finalProject/repository"
	// "github.com/go-playground/validator/v10/translations/id"
)

type ISocialMedia interface {
	Create(request model.SocialMediaCreateRequest, userID string) (model.SocialMediaCreateResponse, error)
	GetAll() ([]model.SocialMediaResponse, error)
	Update(updateReq model.SocialMediaUpdateRequest, SocialID string, userID string) (model.SocialMediaUpdateResponse, error)
	GetOne(socialID string) (model.SocialMediaResponse, error)
}

type SocialMediaService struct {
	SocialMediaRepository repository.ISocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository repository.ISocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		SocialMediaRepository: socialMediaRepository,
	}
}

func (ss *SocialMediaService) Create(request model.SocialMediaCreateRequest, userID string) (model.SocialMediaCreateResponse, error) {
	socialID := helper.GenerateID()

	NewSocial := model.SocialMedia{
		SocialID:       socialID,
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
		UserID:         userID,
	}

	err := ss.SocialMediaRepository.Add(NewSocial)
	if err != nil {
		if err != model.ErrorNotFound {
			return model.SocialMediaCreateResponse{}, err
		}
		return model.SocialMediaCreateResponse{}, model.ErrorNotFound
	}

	response := model.SocialMediaCreateResponse{
		SocialID:       NewSocial.SocialID,
		Name:           NewSocial.Name,
		SocialMediaUrl: NewSocial.SocialMediaUrl,
		UserID:         NewSocial.UserID,
		CreatedAt:      NewSocial.CreatedAt,
	}
	return response, nil
}

func (ss *SocialMediaService) GetAll() ([]model.SocialMediaResponse, error) {
	SocialMediaRes := []model.SocialMediaResponse{}

	res, err := ss.SocialMediaRepository.Get()
	if err != nil {
		return []model.SocialMediaResponse{}, err
	}
	for _, SocialRes := range res {
		SocialMediaRes = append(SocialMediaRes, model.SocialMediaResponse{
			SocialID:       SocialRes.SocialID,
			Name:           SocialRes.Name,
			SocialMediaUrl: SocialRes.SocialMediaUrl,
			UserID:         SocialRes.UserID,
			CreatedAt:      SocialRes.CreatedAt,
			UpdatedAt:      SocialRes.UpdatedAt,
		})
	}

	return SocialMediaRes, nil
}

func (ss *SocialMediaService) Update(updateReq model.SocialMediaUpdateRequest, SocialID string, userID string) (model.SocialMediaUpdateResponse, error) {
	getId, err := ss.SocialMediaRepository.GetOne(SocialID)
	if err != nil {
		if err != model.ErrorNotFound {
			return model.SocialMediaUpdateResponse{}, err
		}
		return model.SocialMediaUpdateResponse{}, model.ErrorNotFound
	}

	if getId.UserID != userID {
		return model.SocialMediaUpdateResponse{}, model.ErrorForbiddenAccess
	}

	SocialUpdate := model.SocialMedia{
		Name:           updateReq.Name,
		SocialMediaUrl: updateReq.SocialMediaUrl,
	}

	res, err := ss.SocialMediaRepository.Update(SocialUpdate, SocialID)

	if err != nil {
		return model.SocialMediaUpdateResponse{}, err
	}

	return model.SocialMediaUpdateResponse{
		SocialID:       res.SocialID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         userID,
		UpdatedAt:      res.UpdatedAt,
	}, nil
}

func (ss *SocialMediaService) GetOne(socialID string) (model.SocialMediaResponse, error) {
	getOne, err := ss.SocialMediaRepository.GetOne(socialID)

	if err != nil {
		if err != model.ErrorNotFound {
			return model.SocialMediaResponse{}, err
		}
		return model.SocialMediaResponse{}, model.ErrorNotFound
	}

	return model.SocialMediaResponse{
		SocialID:       getOne.SocialID,
		Name:           getOne.Name,
		SocialMediaUrl: getOne.SocialMediaUrl,
		UserID:         getOne.UserID,
		CreatedAt:      getOne.CreatedAt,
		UpdatedAt:      getOne.UpdatedAt,
	}, nil
}

func (ss *SocialMediaService) Delete(socialID string, userID string) error {
	getSocialId, err := ss.SocialMediaRepository.GetOne(socialID)

	if err != nil {
		if err != model.ErrorNotFound {
			return err
		}
		return model.ErrorNotFound
	}

	if getSocialId.UserID != userID {
		return model.ErrorForbiddenAccess
	}

	err = ss.SocialMediaRepository.Delete(socialID)

	if err != nil {
		return err
	}

	return nil
}
