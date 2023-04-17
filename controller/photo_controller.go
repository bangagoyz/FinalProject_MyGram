package controller

import (
	"finalProject/model"
	"finalProject/service"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: photoService,
	}
}

// CreatePhoto godoc
//
//		@Summary			Post a Photo on MyGram
//		@Description		Post a Photo on MyGram
//		@Tags				Photo
//		@Accept				json
//		@Produce			json
//		@Param				request body			model.PhotoRequest	true	"Photo request is required"
//		@Success			201		{object}		model.SuccessResponse
//		@Failure			400		{object}		model.FailedResponse
//		@Failure			401		{object}		model.FailedResponse
//		@Failure			500		{object}		model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/photos/create	[post]
func (pc *PhotoController) CreatePhoto(ctx *gin.Context) {
	var request model.PhotoRequest
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

	valid, err := govalidator.ValidateStruct(request)

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

	userID, isExist := ctx.Get("user_id")
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

	photo, err := pc.photoService.Create(request, userID.(string))
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
		Data: photo,
	})

}

// GetAllPhoto godoc
//
//		@Summary			Get All Photo
//		@Description		Show All Photo on MyGram
//		@Tags				Photo
//		@Accept				json
//		@Produce			json
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/photos/get/all	[get]
func (pc *PhotoController) GetAllPhoto(ctx *gin.Context) {
	resp, err := pc.photoService.GetAllPhoto()
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
		Data: resp,
	})
	return
}

// GetOnePhoto godoc
//
//		@Summary			Get One Photo
//		@Description		Show single Photo by input Social Media ID
//		@Tags				Photo
//		@Accept				json
//		@Produce			json
//		@Param				photo_id	path			string 	true		"photo_id"
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			404		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/photos/get/{photo_id}	[get]
func (pc *PhotoController) GetOnePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photo_id")

	response, err := pc.photoService.GetOnePhoto(photoID)
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

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
}

// UpdatePhoto godoc
//
//		@Summary			Update Photo
//		@Description		Update single Photo Title and URL by input Social Media ID
//		@Tags				Photo
//		@Accept				json
//		@Produce			json
//		@Param				photo_id	path			string 	true		"insert your photo id"
//		@Param				request body		model.PhotoRequest	true	"Photo request is required"
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			400		{object}	model.FailedResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			403		{object}	model.FailedResponse
//		@Failure			404		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/photos/update/{photo_id}	[put]
func (pc *PhotoController) PhotoUpdate(ctx *gin.Context) {
	var request model.PhotoRequest
	photoID := ctx.Param("photo_id")
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

	valid, err := govalidator.ValidateStruct(request)
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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	userID, isExist := ctx.Get("user_id")
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

	response, err := pc.photoService.UpdatePhoto(request, userID.(string), photoID)
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
		Data: response,
	})
	return
}

// DeletePhoto godoc
//
//		@Summary			Delete Photo
//		@Description		Delete Photo by input Social Media ID
//		@Tags				Photo
//		@Accept				json
//		@Produce			json
//		@Param				photo_id	path			string 	true		"insert your Photo ID"
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			400		{object}	model.FailedResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			403		{object}	model.FailedResponse
//		@Failure			404		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/photos/delete/{photo_id}	[delete]
func (pc *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photo_id")
	userID, IsExist := ctx.Get("user_id")
	if !IsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: model.ErrorInvalidToken.Err,
		})
		return
	}

	err := pc.photoService.DeletePhoto(photoID, userID.(string))
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
				Error: model.ErrorForbiddenAccess.Err,
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
		Data: "Delete photo success",
	})
	return
}
