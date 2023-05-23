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

type photoHandler struct {
	photoService service.PhotoService
}

func NewPhotoService(photoService service.PhotoService) *photoHandler {
	return &photoHandler{photoService: photoService}
}

func (p *photoHandler) CreatePhoto(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}
	var requestBody dto.CreatePhotoRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdPhoto, err := p.photoService.CreatePhoto(userData, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdPhoto)
}

func (p *photoHandler) GetAllPhotos(ctx *gin.Context) {
	photos, err := p.photoService.GetAllPhotos()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (p *photoHandler) UpdatePhoto(ctx *gin.Context) {
	var requestBody dto.UpdatePhotoRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	photoID := ctx.Param("photoID")
	photoIDUint, err := strconv.ParseUint(photoID, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Photo id should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	updatedPhoto, err2 := p.photoService.UpdatePhoto(uint(photoIDUint), &requestBody)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, updatedPhoto)
}

func (p *photoHandler) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoID")
	photoIDUint, err := strconv.ParseUint(photoID, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Photo id should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	response, err2 := p.photoService.DeletePhoto(uint(photoIDUint))
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
