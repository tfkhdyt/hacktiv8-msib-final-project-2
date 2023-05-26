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

type socialmediaHandler struct {
	socialmediaService service.SocialMediaService
}

func NewSocialMediaService(socialmediaService service.SocialMediaService) *socialmediaHandler {
	return &socialmediaHandler{socialmediaService: socialmediaService}
}

func (s *socialmediaHandler) CreateSocialMedia(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	var requestBody dto.CreateSocialMediaRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdSocialmedia, err := s.socialmediaService.CreateSocialMedia(userData, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdSocialmedia)
}

func (s *socialmediaHandler) GetAllSocialMediasByUserSosmed(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	socialmedias, err := s.socialmediaService.GetAllSocialMediasByUserSosmed(userData.ID)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, socialmedias)
}

func (s *socialmediaHandler) UpdateSocialMedia(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaID")
	socialMediaIDUint, err := strconv.ParseUint(socialMediaID, 10, 32)
	if err != nil {
		errValidation := errs.NewBadRequest("Social Media id should be in unsigned integer")
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	var reqBody dto.UpdateSocialMediaRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errValidation := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	updatedSocialMedia, errUpdate := s.socialmediaService.UpdateSocialMedia(uint(socialMediaIDUint), &reqBody)
	if errUpdate != nil {
		ctx.JSON(errUpdate.StatusCode(), errUpdate)
		return
	}

	ctx.JSON(http.StatusOK, updatedSocialMedia)
}
