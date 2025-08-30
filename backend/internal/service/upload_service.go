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

var queueMap = map[int]string{
	0: "critical",
	1: "critical",
	2: "critical",
	3: "default",
	4: "default",
	5: "low",
	6: "low",
}

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
		final_file := utils.Finalize(uploadId)
		tasks := []*asynq.Task{
			utils.NewVideoTranscode240pTask(final_file, uploadId),
			utils.NewVideoTranscode360pTask(final_file, uploadId),
			utils.NewVideoTranscode480pTask(final_file, uploadId),
			utils.NewVideoTranscode720pTask(final_file, uploadId),
			utils.NewVideoTranscode1080pTask(final_file, uploadId),
			utils.NewVideoTranscode1440pTask(final_file, uploadId),
			utils.NewVideoTranscode2160pTask(final_file, uploadId),
		}

		for priority, t := range tasks {
			q, _ := queueMap[priority]
			if _, err := queueClient.Enqueue(t, asynq.Queue(q)); err != nil {
				log.Printf("failed to enqueue %v", err)
			}
		}
	}
}
