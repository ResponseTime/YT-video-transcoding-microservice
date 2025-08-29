package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/responsetime/video-transcoding-microservice/internal/service"
)

func UploadController() func(c *gin.Context) {
	return func(c *gin.Context) {
		uploadId := c.PostForm("uploadId")
		chunk, _, _ := c.Request.FormFile("chunk")
		part := c.PostForm("part")
		end := c.PostForm("end")
		service.UploadService(uploadId, chunk, part, end)
	}
}
