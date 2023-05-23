package service

import (
	"hacktiv8-msib-final-project-2/dto"
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/repository/comment_repository"
	"hacktiv8-msib-final-project-2/repository/photo_repository"
	"hacktiv8-msib-final-project-2/repository/user_repository"
)

type CommentService interface {
	CreateComment(user *entity.User, payload *dto.CreateCommentRequest) (*dto.CreateCommentResponse, errs.MessageErr)
	GetAllComments() ([]dto.GetAllCommentsResponse, errs.MessageErr)
	UpdateComment(id uint, payload *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr)
	DeleteComment(id uint) (*dto.DeleteCommentResponse, errs.MessageErr)
}

type commentService struct {
	commentRepo comment_repository.CommentRepository
	photoRepo   photo_repository.PhotoRepository
	userRepo    user_repository.UserRepository
}

func NewCommentService(commentRepo comment_repository.CommentRepository, userRepo user_repository.UserRepository) CommentService {
	return &commentService{commentRepo: commentRepo, userRepo: userRepo}
}

func (c *commentService) CreateComment(user *entity.User, payload *dto.CreateCommentRequest) (*dto.CreateCommentResponse, errs.MessageErr) {
	comment := payload.ToEntity()

	createdComment, err := c.commentRepo.CreateComment(user, comment)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateCommentResponse{
		ID:        createdComment.ID,
		Message:   createdComment.Message,
		PhotoID:   createdComment.PhotoID,
		UserID:    createdComment.UserID,
		CreatedAt: createdComment.CreatedAt,
	}

	return response, nil
}

func (c *commentService) GetAllComments() ([]dto.GetAllCommentsResponse, errs.MessageErr) {
	comments, err := c.commentRepo.GetAllComments()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllCommentsResponse{}
	for _, comment := range comments {
		user, err := c.userRepo.GetUserByID(comment.UserID)
		if err != nil {
			return nil, err
		}

		photo, err2 := c.photoRepo.GetPhotoByID(comment.PhotoID)
		if err2 != nil {
			return nil, err2
		}

		response = append(response, dto.GetAllCommentsResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			UpdateAt:  comment.UpdatedAt,
			CreatedAt: comment.CreatedAt,
			User: dto.UserDataWithID{
				ID:       user.ID,
				Email:    user.Email,
				Username: user.Username,
			},
			Photo: dto.PhotoData{
				ID:       photo.ID,
				Title:    photo.Title,
				Caption:  photo.Caption,
				PhotoURL: photo.PhotoURL,
				UserID:   photo.UserID,
			},
		})
	}
	return response, nil
}

func (c *commentService) UpdateComment(id uint, payload *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr) {
	oldComment, err := c.commentRepo.GetCommentByID(id)
	if err != nil {
		return nil, err
	}
	newComment := payload.ToEntity()

	updateComment, err2 := c.commentRepo.UpdateComment(oldComment, newComment)
	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateCommentResponse{
		ID:        updateComment.ID,
		Message:   updateComment.Message,
		PhotoID:   updateComment.PhotoID,
		UserID:    updateComment.UserID,
		UpdatedAt: updateComment.UpdatedAt,
	}

	return response, nil
}

func (c *commentService) DeleteComment(id uint) (*dto.DeleteCommentResponse, errs.MessageErr) {
	if err := c.commentRepo.DeleteComment(id); err != nil {
		return nil, err
	}

	response := &dto.DeleteCommentResponse{
		Message: "Your message has been successfuly deleted",
	}

	return response, nil
}
