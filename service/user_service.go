package service

import (
	"errors"
	"finalProject/helper"
	"finalProject/model"
	"finalProject/repository"

	"gorm.io/gorm"
)

type IUserService interface {
	Register(userRegisterRequest model.UserRegisterRequest) (*model.UserRegisterResponse, error)
	Login(userLoginRequest model.UserLoginRequest) (model.UserLoginResponse, error)
}

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) Register(userRegisterRequest model.UserRegisterRequest) (*model.UserRegisterResponse, error) {
	id := helper.GenerateID()

	hashPassword, err := helper.Hash(userRegisterRequest.Password)
	if err != nil {
		return &model.UserRegisterResponse{}, err
	}

	user := model.User{
		ID:       id,
		Username: userRegisterRequest.Username,
		Email:    userRegisterRequest.Email,
		Password: hashPassword,
		Age:      userRegisterRequest.Age,
	}

	res, err := us.UserRepository.Add(user)

	if err != nil {
		return &model.UserRegisterResponse{}, err
	}

	return &model.UserRegisterResponse{
		ID:        res.ID,
		Username:  res.Username,
		Email:     res.Email,
		Age:       res.Age,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (us *UserService) Login(userLoginRequest model.UserLoginRequest) (model.UserLoginResponse, error) {

	user, err := us.UserRepository.GetByEmail(userLoginRequest.Email)
	if err != nil {
		return model.UserLoginResponse{}, model.ErrorInvalidEmailOrPassword
	}

	if !helper.IsHashValid(user.Password, userLoginRequest.Password) {
		return model.UserLoginResponse{}, model.ErrorInvalidEmailOrPassword
	}

	token, err := helper.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return model.UserLoginResponse{}, model.ErrorInvalidToken
	}

	return model.UserLoginResponse{
		Token: token,
	}, nil

}

func (s *UserService) EmailExists(email string) (bool, error) {
	_, err := s.UserRepository.GetByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return true, err
	}
	if err == nil {
		return true, nil
	}
	return false, nil
}

func (s *UserService) UserNameExists(username string) (bool, error) {
	_, err := s.UserRepository.GetByUsername(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return true, err
	}
	if err == nil {
		return true, nil
	}
	return false, nil
}
