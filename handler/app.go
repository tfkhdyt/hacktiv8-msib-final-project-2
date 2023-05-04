package handler

import (
	"hacktiv8-msib-final-project-2/database"
	"hacktiv8-msib-final-project-2/handler/http_handler"
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

	authService := service.NewAuthService(userRepo)

	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.PUT("/users", authService.Authentication(), userHandler.UpdateUser)
	r.DELETE("/users", authService.Authentication(), userHandler.DeleteUser)

	log.Fatalln(r.Run(":" + PORT))
}
