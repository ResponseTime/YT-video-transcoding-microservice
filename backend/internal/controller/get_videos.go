package controller

import "github.com/gin-gonic/gin"

func GetVideos() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Writer.Write([]byte("Videos From Db"))
	}
}
