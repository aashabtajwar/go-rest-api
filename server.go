package main

import (
	"github.com/aashabtajwar/rest/controller"
	"github.com/aashabtajwar/rest/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	// the server
	server := gin.Default()

	// Context is a struct
	// ctx is a pointer to that struct
	server.GET("/test", func(ctx *gin.Context) {

		// will return a JSON
		// first param is HTTP code
		ctx.JSON(200, gin.H{
			"message": "Okay",
		})
	})
	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8000")
}
