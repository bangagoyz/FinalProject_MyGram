package middleware

import (
	"finalProject/helper"
	"finalProject/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	if auth == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusUnauthorized,
				Message: "Insert Access Token",
			},
			Error: model.ErrorNotAuthorized.Err,
		})
	}

	token := strings.Split(auth, " ")[1]

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			},
			Error: model.ErrorNotAuthorized.Err,
		})
	}

	jwtToken, err := helper.VerifyAccessToken(token)
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

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.FailedResponse{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return

	}

	ctx.Set("user_id", claims["user_id"])

	ctx.Next()
}
