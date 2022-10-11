package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrarkaan/gin-exercises/controller"
	"github.com/jrarkaan/gin-exercises/middleware"
	"github.com/jrarkaan/gin-exercises/service"
	"io"
	"net/http"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	var server *gin.Engine = gin.Default()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasihAuth())

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
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "-1",
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "1",
				"message": "Successfully created a data",
			})
		}

	})

	server.Run(":8585")
}
