package controller

import "github.com/gin-gonic/gin"

func UploadController() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Writer.Write([]byte("Video Upload Controller"))
	}
}
