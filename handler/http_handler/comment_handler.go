package http_handler

import (
	"hacktiv8-msib-final-project-2/dto"
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService service.CommentService
}

func NewCommentService(commentService service.CommentService) *commentHandler {
	return &commentHandler{commentService: commentService}
}

func (c *commentHandler) CreateComment(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	var requestBody dto.CreateCommentRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdComment, err := c.commentService.CreateComment(userData, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdComment)
}

func (c *commentHandler) GetAllCommentsByUserID(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	comments, err := c.commentService.GetAllCommentsByUserID(userData.ID)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (c *commentHandler) UpdateComment(ctx *gin.Context) {
	commentID := ctx.Param("commentID")
	commentIDUint, err := strconv.ParseUint(commentID, 10, 32)
	if err != nil {
		errValidation := errs.NewBadRequest("Comment id should be in unsigned integer")
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	var reqBody dto.UpdateCommentRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errValidation := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	updatedComment, errUpdate := c.commentService.UpdateComment(uint(commentIDUint), &reqBody)
	if errUpdate != nil {
		ctx.JSON(errUpdate.StatusCode(), errUpdate)
		return
	}

	ctx.JSON(http.StatusOK, updatedComment)
}
