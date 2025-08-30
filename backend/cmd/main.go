package main

import (
	"github.com/responsetime/video-transcoding-microservice/internal/router"
	"github.com/responsetime/video-transcoding-microservice/internal/utils"
)

func main() {
	utils.InitRedis()
	router := router.GenRouter()
	router.Run(":5000")
}
