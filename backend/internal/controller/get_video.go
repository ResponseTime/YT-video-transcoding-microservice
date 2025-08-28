package controller

import "github.com/gin-gonic/gin"

func GetVideoController() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Writer.Write([]byte("Get Video Controller"))
	}
}
