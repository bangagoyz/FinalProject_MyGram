package controller

import (
	"finalProject/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: model.HomeInformationResponse{
			About:    "Final Project MyGram - Hacktiv8 DTSFGA : Scallable Web Service With Golang",
			Name:     "Yoga Budi Permana Putra",
			Github:   "https://github.com/bangagoyz",
			LinkedIn: "https://www.linkedin.com/in/yoga-budi-permana-putra-b62247256/",
		},
	})
}
