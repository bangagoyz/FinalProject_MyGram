package helper

import (
	"finalProject/model"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateUserRegisterRequest(request model.UserRegisterRequest) []error {
	validate := validator.New()

	err := validate.Struct(request)
	if err != nil {
		var errors []error
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				errors = append(errors, fmt.Errorf("%s is required", e.Field()))
			case "email":
				errors = append(errors, fmt.Errorf("%s must be a valid email address", e.Field()))
			case "min":
				errors = append(errors, fmt.Errorf("%s must be at least %s characters", e.Field(), e.Param()))
			case "gt":
				errors = append(errors, fmt.Errorf("%s must be greater than %s", e.Field(), e.Param()))
			default:
				errors = append(errors, fmt.Errorf("%s is invalid", e.Field()))
			}
		}
		return errors
	}

	return nil
}

func ValidateCommentRequest(request model.CommentCreateRequest) []error {
	validate := validator.New()

	err := validate.Struct(request)
	if err != nil {
		var errors []error
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				errors = append(errors, fmt.Errorf("%s is required", e.Field()))
			default:
				errors = append(errors, fmt.Errorf("%s is invalid", e.Field()))
			}
		}
		return errors
	}

	return nil
}

func ValidatePhotoRequest(request model.PhotoRequest) []error {
	validate := validator.New()

	err := validate.Struct(request)
	if err != nil {
		var errors []error
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				errors = append(errors, fmt.Errorf("%s is required", e.Field()))
			default:
				errors = append(errors, fmt.Errorf("%s is invalid", e.Field()))
			}
		}
		return errors
	}

	return nil
}
