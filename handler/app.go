package handler

import (
	"hacktiv8-msib-final-project-2/database"
	"hacktiv8-msib-final-project-2/handler/http_handler"
	commment_pg "hacktiv8-msib-final-project-2/repository/comment_repository/comment_pg"
	"hacktiv8-msib-final-project-2/repository/photo_repository/photo_pg"
	"hacktiv8-msib-final-project-2/repository/user_repository/user_pg"
	"hacktiv8-msib-final-project-2/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	commentRepo := commment_pg.NewCommentPG(db)
	commentService := service.NewCommentService(commentRepo, userRepo)
	commentHandler := http_handler.NewCommentService(commentService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo)

	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.PUT("/users", authService.Authentication(), userHandler.UpdateUser)
	r.DELETE("/users", authService.Authentication(), userHandler.DeleteUser)

	r.POST("/photos", authService.Authentication(), photoHandler.CreatePhoto)
	r.GET("/photos", authService.Authentication(), photoHandler.GetAllPhotos)
	r.PUT("/photos/:photoID", authService.Authentication(), authService.PhotosAuthorization(), photoHandler.UpdatePhoto)
	r.DELETE("/photos/:photoID", authService.Authentication(), authService.PhotosAuthorization(), photoHandler.DeletePhoto)

	r.POST("/comments", authService.Authentication(), authService.CommentsAuthorization(), commentHandler.CreateComment)

	log.Fatalln(r.Run(":" + PORT))
}
