package http_handler

import (
	"hacktiv8-msib-final-project-2/dto"
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/service"
	"net/http"

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
