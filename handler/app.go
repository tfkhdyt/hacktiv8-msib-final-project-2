package handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"hacktiv8-msib-final-project-2/database"
	"hacktiv8-msib-final-project-2/handler/http_handler"
	"hacktiv8-msib-final-project-2/repository/comment_repository/comment_pg"
	"hacktiv8-msib-final-project-2/repository/photo_repository/photo_pg"
	"hacktiv8-msib-final-project-2/repository/socialmedia_repository/socialmedia_pg"
	"hacktiv8-msib-final-project-2/repository/user_repository/user_pg"
	"hacktiv8-msib-final-project-2/service"
)

var PORT = os.Getenv("PORT")

func StartApp() {
	db := database.GetPostgresInstance()

	if PORT == "" {
		PORT = "8080"
	}
	r := gin.Default()

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := http_handler.NewUserHandler(userService)

	photoRepo := photo_pg.NewPhotoPG(db)
	photoService := service.NewPhotoService(photoRepo, userRepo)
	photoHandler := http_handler.NewPhotoService(photoService)

	commentRepo := comment_pg.NewCommentPG(db)
	commentService := service.NewCommentService(commentRepo, photoRepo, userRepo)
	commentHandler := http_handler.NewCommentService(commentService)

	socialmediaRepo := socialmedia_pg.NewSocialMediaPG(db)
	socialmediaService := service.NewSocialMediaService(socialmediaRepo, userRepo)
	socialmediaHandler := http_handler.NewSocialMediaService(socialmediaService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, socialmediaRepo)

	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.PUT("/users", authService.Authentication(), userHandler.UpdateUser)
	r.DELETE("/users", authService.Authentication(), userHandler.DeleteUser)

	r.POST("/photos", authService.Authentication(), photoHandler.CreatePhoto)
	r.GET("/photos", authService.Authentication(), photoHandler.GetAllPhotos)
	r.PUT("/photos/:photoID", authService.Authentication(), authService.PhotosAuthorization(), photoHandler.UpdatePhoto)
	r.DELETE("/photos/:photoID", authService.Authentication(), authService.PhotosAuthorization(), photoHandler.DeletePhoto)

	r.POST("/comments", authService.Authentication(), commentHandler.CreateComment)
	r.GET("/comments", authService.Authentication(), commentHandler.GetAllCommentsByUserID)
	r.PUT("/comments/:commentID", authService.Authentication(), authService.CommentsAuthorization(), commentHandler.UpdateComment)
	r.DELETE("/comments/:commentID", authService.Authentication(), authService.CommentsAuthorization(), commentHandler.DeleteComment)

	r.POST("/socialmedias", authService.Authentication(), socialmediaHandler.CreateSocialMedia)
	r.GET("/socialmedias", authService.Authentication(), socialmediaHandler.GetAllSocialMediasByUserSosmed)
	r.PUT("/socialmedias/:socialMediaID", authService.Authentication(), authService.SocialmediasAuthorization(), socialmediaHandler.UpdateSocialMedia)
	r.DELETE("/socialmedias/:socialMediaID", authService.Authentication(), authService.SocialmediasAuthorization(), socialmediaHandler.DeleteSocialMedia)

	log.Fatalln(r.Run(":" + PORT))
}
