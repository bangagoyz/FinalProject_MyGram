package router

import (
	"finalProject/controller"
	"finalProject/middleware"
	"finalProject/repository"
	"finalProject/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "finalProject/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title							Mygram API
// @version							1.0
// @description						Final Project for Scalable Web Service with Golang DTS-FGA
// @description						by Yoga Budi Permana Putra. \n
// @description						NOTE : input Authorize button format : bearer YOURACCESSTOKEN
// @accept							json
// @produce							json
// @securityDefinitions.apikey		Bearer
// @in								header
// @name							Authorization
func StartApp(router *gin.Engine, db *gorm.DB) {
	commentRepository := repository.NewCommentRepository(db)
	userRepository := repository.NewUserRepository(db)
	photoRepository := repository.NewPhotoRepository(db)
	SocialMediaRepository := repository.NewSocialMediaRepository(db)

	photoService := service.NewPhotoService(photoRepository, commentRepository)
	photoController := controller.NewPhotoController(*photoService)

	commentService := service.NewCommentService(commentRepository, photoRepository)
	commentController := controller.NewCommentController(*commentService)

	userService := service.NewUserService(*userRepository)
	userController := controller.NewUserController(*userService)

	SocialMediaService := service.NewSocialMediaService(SocialMediaRepository)
	SocialMediaController := controller.NewSocialMediaController(*SocialMediaService)

	base := router.Group("/mygram")
	{
		user := base.Group("/user")
		{
			user.POST("/register", userController.Register)
			user.POST("/login", userController.Login)
		}
		withAuth := base.Group("/photos", middleware.AuthMiddleware)
		{
			withAuth.POST("/create", photoController.CreatePhoto)
			withAuth.GET("/get/all", photoController.GetAllPhoto)
			withAuth.GET("/get/:photo_id", photoController.GetOnePhoto)
			withAuth.PUT("/update/:photo_id", photoController.PhotoUpdate)
			withAuth.DELETE("/delete/:photo_id", photoController.DeletePhoto)
		}
		commentAuth := base.Group("/comments", middleware.AuthMiddleware)
		{
			commentAuth.POST("/:photo_id", commentController.CreateComment)
			commentAuth.GET("/get/all", commentController.GetAllComment)
			commentAuth.GET("/get/:comment_id", commentController.GetOneComment)
			commentAuth.PUT("/update/:comment_id", commentController.UpdateComment)
			commentAuth.DELETE("/delete/:comment_id", commentController.DeleteComment)
		}
		socialAuth := base.Group("/social_media", middleware.AuthMiddleware)
		{
			socialAuth.POST("/", SocialMediaController.CreateSocialMedia)
			socialAuth.GET("/get/all", SocialMediaController.GetAllSocialMedia)
			socialAuth.PUT("/update/:social_id", SocialMediaController.UpdateSocialMedia)
			socialAuth.GET("/get/:social_id", SocialMediaController.GetOneSocial)
			socialAuth.DELETE("/delete/:social_id", SocialMediaController.DeleteSocial)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
