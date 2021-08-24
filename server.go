package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wilker/go-with-gin/controller"
	"github.com/wilker/go-with-gin/service"
)

var (
	videoService    = service.New()
	videoController = controller.New(videoService)
)

func main(){
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context){
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context){
		ctx.JSON(201, videoController.Save(ctx))
	})

	server.Run(":8000")
}
