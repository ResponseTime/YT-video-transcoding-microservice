package main

import (
	"github.com/responsetime/video-transcoding-microservice/internal/router"
	"github.com/responsetime/video-transcoding-microservice/internal/utils"
)

func main() {
	utils.InitRedis()
	worker, mux := utils.Queue()
	router := router.GenRouter()
	worker.Run(mux)
	router.Run(":5000")
}
