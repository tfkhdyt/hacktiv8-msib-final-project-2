package service

import (
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/photo_repository"
	"hacktiv8-msib-final-project-2/repository/user_repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	PhotosAuthorization() gin.HandlerFunc
	// CommentsAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo  user_repository.UserRepository
	photoRepo photo_repository.PhotoRepository
	// commentRepo comment_repository.CommentRepository
}

func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository) AuthService {
	return &authService{userRepo: userRepo, photoRepo: photoRepo}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		if err := user.ValidateToken(bearerToken); err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		result, err := a.userRepo.GetUserByID(user.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		ctx.Set("userData", result)
		ctx.Next()
	}
}

func (a *authService) PhotosAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		photoID := ctx.Param("photoID")
		photoIDUint, err := strconv.ParseUint(photoID, 10, 32)
		if err != nil {
			newError := errs.NewBadRequest("Photo id should be an unsigned integer")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		photo, err2 := a.photoRepo.GetPhotoByID(uint(photoIDUint))
		if err2 != nil {
			ctx.AbortWithStatusJSON(err2.StatusCode(), err2)
			return
		}

		if photo.UserID != userData.ID {
			newError := errs.NewUnauthorized("You're not authorized to modify this photo")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}

// func (a *authService) CommentsAuthorization() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		userData, ok := ctx.MustGet("userData").(*entity.User)
// 		if !ok {
// 			newError := errs.NewBadRequest("Failed to get user data")
// 			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
// 			return
// 		}
//
// 		commentID := ctx.Param("commentID")
// 		commentIDUint, err := strconv.ParseUint(commentID, 10, 32)
// 		if err != nil {
// 			newError := errs.NewBadRequest("Comment id should be an unsigned integer")
// 			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
// 			return
// 		}
//
// 		comment, err2 := a.commentRepo.GetCommentByID(uint(commentIDUint))
// 		if err2 != nil {
// 			ctx.AbortWithStatusJSON(err2.StatusCode(), err2)
// 			return
// 		}
//
// 		if comment.UserID != userData.ID {
// 			newError := errs.NewUnauthorized("You're not authorized to modify this comment")
// 			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
// 			return
// 		}
//
// 		ctx.Next()
// 	}
// }
