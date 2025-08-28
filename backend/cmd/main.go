package main

import "github.com/responsetime/video-transcoding-microservice/internal/router"

func main() {
	router := router.GenRouter()
	router.Run(":3000")
}
