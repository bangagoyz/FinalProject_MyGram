package controller

import (
	"finalProject/model"
	"finalProject/service"
	"net/http"

	Valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService service.UserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

// Register godoc
//
//		@Summary			Register User
//		@Description		Register MyGram User
//		@Tags				User
//		@Accept				json
//		@Produce			json
//		@Param				request body			model.UserRegisterRequest	true	"minimum age to register is 8 years old. || password minimum is 6 character."
//		@Success			201		{object}		model.SuccessResponse
//		@Failure			400		{object}		model.FailedResponse
//		@Failure			500		{object}		model.FailedResponse
//	 @Router				/mygram/user/register	[post]
func (uc *UserController) Register(ctx *gin.Context) {
	var request model.UserRegisterRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	valid, err := Valid.ValidateStruct(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	response, err := uc.UserService.Register(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return

	}
	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
	return

}

// Register godoc
//
//		@Summary			Login User
//		@Description		Sign in MyGram User to access all feature. NOTE : to input access token to Authorize button, please write with format: bearer YourTokenAccess || Token will be expired in 1 hours
//		@Tags				User
//		@Accept				json
//		@Produce			json
//		@Param				request body			model.UserLoginRequest	true	"User Login Request is required"
//		@Success			201		{object}		model.SuccessResponse
//		@Failure			400		{object}		model.FailedResponse
//		@Failure			500		{object}		model.FailedResponse
//	 @Router				/mygram/user/login	[post]
func (uc *UserController) Login(ctx *gin.Context) {
	var request model.UserLoginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := Valid.ValidateStruct(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	response, err := uc.UserService.Login(request)
	if err != nil {
		if err == model.ErrorInvalidEmailOrPassword {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.FailedResponse{
				Meta: model.Meta{
					Code:    http.StatusUnauthorized,
					Message: http.StatusText(http.StatusUnauthorized),
				},
				Error: err.Error(),
			})
			return
		} else if err == model.ErrorInvalidToken {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.FailedResponse{
				Meta: model.Meta{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
				},
				Error: model.ErrorInvalidToken.Err,
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
}
