package controller

import (
	"github.com/aashabtajwar/rest/entity"
	"github.com/aashabtajwar/rest/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	// finds all videos (a slice of videos) of type video (from entity package)
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video

}
