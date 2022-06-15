package controller

import (
	"github.com/aashabtajwar/rest/entity"
	"github.com/aashabtajwar/rest/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{service: }
}
