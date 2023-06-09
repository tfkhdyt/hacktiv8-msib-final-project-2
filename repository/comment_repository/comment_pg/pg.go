package comment_pg

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/comment_repository"
)

type commentPG struct {
	db *gorm.DB
}

func NewCommentPG(db *gorm.DB) comment_repository.CommentRepository {
	return &commentPG{db: db}
}

func (c *commentPG) CreateComment(user *entity.User, comment *entity.Comment) (*entity.Comment, errs.MessageErr) {
	if err := c.db.Model(user).Association("Comments").Append(comment); err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to create new comment")
	}

	return comment, nil
}

func (c *commentPG) GetAllCommentsByUserID(userID uint) ([]entity.Comment, errs.MessageErr) {
	var comments []entity.Comment
	if err := c.db.Find(&comments, "user_id = ?", userID).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to get all comment")
	}

	return comments, nil
}

func (c *commentPG) GetCommentByID(id uint) (*entity.Comment, errs.MessageErr) {
	var comment entity.Comment
	if err := c.db.First(&comment, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Comment with id %d is not found", id))
	}

	return &comment, nil
}

func (c *commentPG) UpdateComment(oldComment *entity.Comment, newComment *entity.Comment) (*entity.Comment, errs.MessageErr) {
	if err := c.db.Model(oldComment).Updates(newComment).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update comment with id %d", oldComment.ID))
	}

	return oldComment, nil
}

func (c *commentPG) DeleteComment(id uint) errs.MessageErr {
	if err := c.db.Delete(&entity.Comment{}, id).Error; err != nil {
		log.Println("Error:", err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete comment with id %d", id))
	}

	return nil
}
