package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrarkaan/gin-exercises/controller"
	"github.com/jrarkaan/gin-exercises/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	var server *gin.Engine = gin.Default()

	server.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status":  1,
			"message": "Get The Server!",
		})
	})

	server.GET("/api/videos", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status":  1,
			"message": "Get The Data!",
			"data":    videoController.FindAll(),
		})
	})

	server.POST("/api/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8585")
}
