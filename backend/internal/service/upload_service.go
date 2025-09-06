package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/responsetime/video-transcoding-microservice/internal/utils"
)

func UploadService(uploadId string, chunk multipart.File, part string, end string) {
	redisInstance := utils.GetRedisInstance()
	queueClient := utils.QueueClient()
	ctx := context.Background()
	var file *os.File
	partn, _ := strconv.Atoi(part)
	file, err := os.Create(fmt.Sprintf("./internal/temp/%s-part-%d.mp4", uploadId, partn))
	redisInstance.ZAdd(ctx, uploadId, redis.Z{
		Score:  float64(partn),
		Member: fmt.Sprintf("./internal/temp/%s-part-%d.mp4", uploadId, partn),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := io.Copy(file, chunk); err != nil {
		log.Fatal(err)
	}
	if end == "True" {
		FntTask := utils.NewFnTTask(uploadId)
		queueClient.Enqueue(FntTask, asynq.Queue("default"))
	}
}
