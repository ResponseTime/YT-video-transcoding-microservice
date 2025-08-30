package utils

import (
	"context"
	"encoding/json"
	"runtime"

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
	UploadID string `json:"upload_id"`
	Path     string `json:"path"`
}

func Queue() (*asynq.Server, *asynq.ServeMux) {
	worker := asynq.NewServer(redisConn, asynq.Config{
		Concurrency: runtime.NumCPU() - 1,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeVideoTranscode240p, HandleVideoTranscoding240p)
	mux.HandleFunc(TypeVideoTranscode360p, HandleVideoTranscoding360p)
	mux.HandleFunc(TypeVideoTranscode480p, HandleVideoTranscoding480p)
	mux.HandleFunc(TypeVideoTranscode720p, HandleVideoTranscoding720p)
	mux.HandleFunc(TypeVideoTranscode1080p, HandleVideoTranscoding1080p)
	mux.HandleFunc(TypeVideoTranscode1440p, HandleVideoTranscoding1440p)
	mux.HandleFunc(TypeVideoTranscode2160p, HandleVideoTranscoding2160p)
	return worker, mux
}

func QueueClient() *asynq.Client {
	return asynq.NewClient(redisConn)
}

// -------------------- 240p --------------------
func NewVideoTranscode240pTask(videoPath string, uploadId string) *asynq.Task {
	payload, _ := json.Marshal(map[string]interface{}{"upload_id": uploadId, "path": videoPath})
	return asynq.NewTask(TypeVideoTranscode240p, payload)
}

func HandleVideoTranscoding240p(c context.Context, t *asynq.Task) error {
	var p payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// TODO: actual transcoding logic for 240p
	return nil
}

// -------------------- 360p --------------------
func NewVideoTranscode360pTask(videoPath string, uploadId string) *asynq.Task {
	payload, _ := json.Marshal(map[string]interface{}{"upload_id": uploadId, "path": videoPath})
	return asynq.NewTask(TypeVideoTranscode360p, payload)
}

func HandleVideoTranscoding360p(c context.Context, t *asynq.Task) error {
	var p payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// TODO: actual transcoding logic for 360p
	return nil
}

// -------------------- 480p --------------------
func NewVideoTranscode480pTask(videoPath string, uploadId string) *asynq.Task {
	payload, _ := json.Marshal(map[string]interface{}{"upload_id": uploadId, "path": videoPath})
	return asynq.NewTask(TypeVideoTranscode480p, payload)
}

func HandleVideoTranscoding480p(c context.Context, t *asynq.Task) error {
	var p payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// TODO: actual transcoding logic for 480p
	return nil
}

// -------------------- 720p --------------------
func NewVideoTranscode720pTask(videoPath string, uploadId string) *asynq.Task {
	payload, _ := json.Marshal(map[string]interface{}{"upload_id": uploadId, "path": videoPath})
	return asynq.NewTask(TypeVideoTranscode720p, payload)
}

func HandleVideoTranscoding720p(c context.Context, t *asynq.Task) error {
	var p payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// TODO: actual transcoding logic for 720p
	return nil
}

// -------------------- 1080p --------------------
func NewVideoTranscode1080pTask(videoPath string, uploadId string) *asynq.Task {
	payload, _ := json.Marshal(map[string]interface{}{"upload_id": uploadId, "path": videoPath})
	return asynq.NewTask(TypeVideoTranscode1080p, payload)
}

func HandleVideoTranscoding1080p(c context.Context, t *asynq.Task) error {
	var p payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// TODO: actual transcoding logic for 1080p
	return nil
}

// -------------------- 1440p --------------------
func NewVideoTranscode1440pTask(videoPath string, uploadId string) *asynq.Task {
	payload, _ := json.Marshal(map[string]interface{}{"upload_id": uploadId, "path": videoPath})
	return asynq.NewTask(TypeVideoTranscode1440p, payload)
}

func HandleVideoTranscoding1440p(c context.Context, t *asynq.Task) error {
	var p payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// TODO: actual transcoding logic for 1440p
	return nil
}

// -------------------- 2160p --------------------
func NewVideoTranscode2160pTask(videoPath string, uploadId string) *asynq.Task {
	payload, _ := json.Marshal(map[string]interface{}{"upload_id": uploadId, "path": videoPath})
	return asynq.NewTask(TypeVideoTranscode2160p, payload)
}

func HandleVideoTranscoding2160p(c context.Context, t *asynq.Task) error {
	var p payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// TODO: actual transcoding logic for 2160p (4K)
	return nil
}
