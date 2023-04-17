package controller

import (
	"finalProject/model"
	"finalProject/service"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{
		CommentService: commentService,
	}
}

// CreateComment godoc
//
//		@Summary			Post Comment on Photo
//		@Description		Post Comment on Photo
//		@Tags				Comment
//		@Accept				json
//		@Produce			json
//		@Param				photo_id	path		string	true	"Photo ID"
//		@Param				request body			model.CommentCreateRequest	true	"Comment Create Request is required"
//		@Success			201		{object}		model.SuccessResponse
//		@Failure			400		{object}		model.FailedResponse
//		@Failure			401		{object}		model.FailedResponse
//		@Failure			404		{object}		model.FailedResponse
//		@Failure			401		{object}		model.FailedResponse
//		@Security			Bearer
//	 	@Router				/mygram/comments/{photo_id}	[post]
func (cc *CommentController) CreateComment(ctx *gin.Context) {
	var request model.CommentCreateRequest
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
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

	comment, err := cc.CommentService.CreateComment(request, userID.(string), photoID)
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
		Data: comment,
	})
}

// GetAllComment godoc
//
//		@Summary			Get All Comment
//		@Description		Show All Comment on MyGram
//		@Tags				Comment
//		@Accept				json
//		@Produce			json
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/comments/get/all	[get]
func (cc *CommentController) GetAllComment(ctx *gin.Context) {
	allComment, err := cc.CommentService.GetAll()
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
		Data: allComment,
	})
	return
}

// GetOneComment godoc
//
//		@Summary			Get One Comment
//		@Description		Show single Comment by input Social Media ID
//		@Tags				Comment
//		@Accept				json
//		@Produce			json
//		@Param				comment_id	path			string 	true		"insert your Comment ID"
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			404		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/comments/get/{comment_id}	[get]
func (cc *CommentController) GetOneComment(ctx *gin.Context) {
	CommentID := ctx.Param("comment_id")
	OneComment, err := cc.CommentService.GetOne(CommentID)

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
		Data: OneComment,
	})
	return
}

// CommentPhoto godoc
//
//		@Summary			Update Comment
//		@Description		Update single Comment by input Social Media ID
//		@Tags				Comment
//		@Accept				json
//		@Produce			json
//		@Param				comment_id	path			string 	true		"insert your Comment ID"
//		@Param				request body		model.CommentUpdateRequest	true	"Comment Update Request is required"
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			400		{object}	model.FailedResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			403		{object}	model.FailedResponse
//		@Failure			404		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/comments/update/{comment_id}	[put]
func (cc *CommentController) UpdateComment(ctx *gin.Context) {
	var UpComment model.CommentUpdateRequest
	CommentId := ctx.Param("comment_id")

	if err := ctx.ShouldBindJSON(&UpComment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := govalidator.ValidateStruct(UpComment)
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

	res, err := cc.CommentService.Update(UpComment, CommentId, userId.(string))
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
		Data: res,
	})
	return
}

// DeleteComment godoc
//
//		@Summary			Delete Comment
//		@Description		Delete Comment by input Social Media ID
//		@Tags				Comment
//		@Accept				json
//		@Produce			json
//		@Param				comment_id	path			string 	true		"insert your Comment ID"
//		@Success			201		{object}	model.SuccessResponse
//		@Failure			400		{object}	model.FailedResponse
//		@Failure			401		{object}	model.FailedResponse
//		@Failure			403		{object}	model.FailedResponse
//		@Failure			404		{object}	model.FailedResponse
//		@Failure			500		{object}	model.FailedResponse
//		@Security			Bearer
//	 @Router				/mygram/comments/delete/{comment_id}	[delete]
func (cc *CommentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("comment_id")
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

	err := cc.CommentService.Delete(commentID, UserID.(string))
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
		Data: "Delete Comment Success!",
	})
}
