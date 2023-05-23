package comment_repository

import (
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
)

type CommentRepository interface {
	CreateComment(user *entity.User, comment *entity.Comment) (*entity.Comment, errs.MessageErr)
	GetAllCommentsByUserID(userID uint) ([]entity.Comment, errs.MessageErr)
	GetCommentByID(id uint) (*entity.Comment, errs.MessageErr)
	UpdateComment(oldComment *entity.Comment, newComment *entity.Comment) (*entity.Comment, errs.MessageErr)
	// DeleteComment(id uint) errs.MessageErr
}
