package utils

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypeVideoTranscode240p  = "transcode:240p"
	TypeVideoTranscode360p  = "transcode:360p"
	TypeVideoTranscode480p  = "transcode:480p"
	TypeVideoTranscode720p  = "transcode:720p"
	TypeVideoTranscode1080p = "transcode:1080p"
	TypeVideoTranscode1440p = "transcode:1440p"
	TypeVideoTranscode2160p = "transcode:2160p"
)

var redisConn = asynq.RedisClientOpt{Addr: "localhost:6379"}

type payload struct {
	uploadId string
	path     string
}

func Queue() (*asynq.Server, *asynq.ServeMux) {
	worker := asynq.NewServer(redisConn, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeVideoTranscode240p, HandleVideoTranscoding240p)
	return worker, mux
}

func QueueClient() *asynq.Client {
	return asynq.NewClient(redisConn)
}

func NewVideoTranscode240pTask(videoPath string, uploadId string) *asynq.Task {
	payload, _ := json.Marshal(map[string]interface{}{"upload_id": uploadId, "path": videoPath})
	return asynq.NewTask(TypeVideoTranscode240p, payload)
}

func HandleVideoTranscoding240p(c context.Context, t *asynq.Task) error {
	var p payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	return nil
}
