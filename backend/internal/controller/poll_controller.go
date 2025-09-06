package controller

import (
	"github.com/gin-gonic/gin"
)

func PollController() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Writer.Write([]byte("Poll Controller"))
	}
}
