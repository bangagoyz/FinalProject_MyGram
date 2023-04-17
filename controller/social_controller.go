package controller

import (
	"finalProject/model"
	"finalProject/service"

	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		SocialMediaService: socialMediaService,
	}
}

// CreateSocialMedia godoc
//
//		@Summary			Create Social Media Account
//		@Description		Add new social Media Account
//		@Tags				Social Media
//		@Accept				json
//		@Produce			json
//		@Param				request body			model.SocialMediaCreateRequest	true	"Social Media request is required"
//		@Success			201		{object}		model.SuccessResponse
//		@Failure			400		{object}		model.FailedResponse
//		@Failure			401		{object}		model.FailedResponse
//		@Failure			500		{object}		model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/social_media/	[post]
func (sc *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var request model.SocialMediaCreateRequest
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

	valid, err := valid.ValidateStruct(request)

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

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: model.ErrorInvalidToken.Err,
		})
		return
	}

	res, err := sc.SocialMediaService.Create(request, userId.(string))
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

	ctx.JSON(http.StatusCreated, model.SuccessResponse{
		Meta: model.Meta{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		},
		Data: res,
	})
	return
}

// GetAllSocialMedia godoc
//
//		@Summary			Get All Social Media
//		@Description		Show All Social Media Account
//		@Tags				Social Media
//		@Accept				json
//		@Produce			json
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/social_media/get/all	[get]
func (sc *SocialMediaController) GetAllSocialMedia(ctx *gin.Context) {
	AllSocMed, err := sc.SocialMediaService.GetAll()
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
		Data: AllSocMed,
	})
	return
}

// GetOneSocialMedia godoc
//
//		@Summary			Get One Social Media Account
//		@Description		Show single Social Media Account by input Social Media ID
//		@Tags				Social Media
//		@Accept				json
//		@Produce			json
//		@Param				social_id	path			string 	true		"insert Social Media ID"
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			404		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/social_media/get/{social_id}	[get]
func (sc *SocialMediaController) GetOneSocial(ctx *gin.Context) {
	SocialId := ctx.Param("social_id")
	socMed, err := sc.SocialMediaService.GetOne(SocialId)

	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.FailedResponse{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: socMed,
	})
}

// UpdateSocialMedia godoc
//
//		@Summary			Update Social Media Account
//		@Description		Update single Social Media Account by input Social Media ID
//		@Tags				Social Media
//		@Accept				json
//		@Produce			json
//		@Param				social_id	path			string 	true		"insert Social Media ID"
//		@Param				request body		model.SocialMediaUpdateRequest	true	"Social Media Update request is required"
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			400		{object}	model.FailedResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			403		{object}	model.FailedResponse
//		@Failure			404		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/social_media/update/{social_id}	[put]
func (sc *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	var UpSocial model.SocialMediaUpdateRequest
	SocialId := ctx.Param("social_id")

	if err := ctx.ShouldBindJSON(&UpSocial); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(UpSocial)
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

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: model.ErrorInvalidToken.Err,
		})
		return
	}

	res, err := sc.SocialMediaService.Update(UpSocial, SocialId, userId.(string))

	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.FailedResponse{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: err.Error(),
			})
			return
		} else if err == model.ErrorForbiddenAccess {
			ctx.AbortWithStatusJSON(http.StatusForbidden, model.FailedResponse{
				Meta: model.Meta{
					Code:    http.StatusForbidden,
					Message: http.StatusText(http.StatusForbidden),
				},
				Error: err.Error(),
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
		Data: res,
	})
	return
}

// DeleteSocialMedia godoc
//
//		@Summary			Delete Social Media Account
//		@Description		Delete single Social Media Account by input Social Media ID
//		@Tags				Social Media
//		@Accept				json
//		@Produce			json
//		@Param				social_id	path			string 	true		"insert Social Media ID"
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			400		{object}	model.FailedResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			403		{object}	model.FailedResponse
//		@Failure			404		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/social_media/delete/{social_id}	[delete]
func (sc *SocialMediaController) DeleteSocial(ctx *gin.Context) {
	SocialId := ctx.Param("social_id")
	UserID, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: model.ErrorInvalidToken.Err,
		})
		return
	}

	err := sc.SocialMediaService.Delete(SocialId, UserID.(string))
	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.FailedResponse{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: model.ErrorNotFound.Err,
			})
			return
		} else if err == model.ErrorForbiddenAccess {
			ctx.AbortWithStatusJSON(http.StatusForbidden, model.FailedResponse{
				Meta: model.Meta{
					Code:    http.StatusForbidden,
					Message: http.StatusText(http.StatusForbidden),
				},
				Error: err.Error(),
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
		Data: "Delete social media success",
	})
	return

}
