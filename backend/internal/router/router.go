package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/responsetime/video-transcoding-microservice/internal/controller"
)

func GenRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/health-check", func(c *gin.Context) {
		c.Writer.Write([]byte("Healthy"))
	})
	router.POST("/upload-video", controller.UploadController())
	router.GET("/poll-video-metadata/:uploadId", controller.PollController())
	router.GET("/get-videos", controller.GetVideoController())
	router.GET("/video/:uploadId/:resolution", controller.GetVideoController())
	return router
}
