package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
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
	TypeFnT                 = "FnT"
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

var redisConn = asynq.RedisClientOpt{Addr: "localhost:6379"}
var db = DBConn()

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
	mux.HandleFunc(TypeFnT, HandleFnTTask)
	return worker, mux
}

func QueueClient() *asynq.Client {
	return asynq.NewClient(redisConn)
}

// Finalize file upload & create video transcoding jobs

func NewFnTTask(uploadId string) *asynq.Task {
	payload, _ := json.Marshal(map[string]interface{}{"upload_id": uploadId})
	return asynq.NewTask(TypeFnT, payload)
}

func HandleFnTTask(c context.Context, t *asynq.Task) error {
	var p payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	final_file := Finalize(p.UploadID)
	tasks := []*asynq.Task{
		NewVideoTranscode240pTask(final_file, p.UploadID),
		NewVideoTranscode360pTask(final_file, p.UploadID),
		NewVideoTranscode480pTask(final_file, p.UploadID),
		NewVideoTranscode720pTask(final_file, p.UploadID),
		NewVideoTranscode1080pTask(final_file, p.UploadID),
		NewVideoTranscode1440pTask(final_file, p.UploadID),
		NewVideoTranscode2160pTask(final_file, p.UploadID),
	}

	for priority, t := range tasks {
		q, _ := queueMap[priority]
		if _, err := QueueClient().Enqueue(t, asynq.Queue(q)); err != nil {
			log.Printf("failed to enqueue %v", err)
		}
	}
	return nil
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

	// 240p
	output240p := path.Join(fmt.Sprintf("./internal/temp/%s-240p.mp4", p.UploadID))
	cmd240p := exec.Command(
		"ffmpeg",
		"-i", p.Path,
		"-vf", "scale=-2:240",
		"-c:v", "libx264",
		"-preset", "fast",
		"-crf", "23",
		"-c:a", "aac",
		"-b:a", "128k",
		"-y",
		output240p,
	)
	cmd240p.Stdout = os.Stdout
	cmd240p.Stderr = os.Stderr
	cmd240p.Run()
	cmd240p.Wait()
	// db.Exec()
	// db.Acquire()
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
	output360p := path.Join(fmt.Sprintf("./internal/temp/%s-360p.mp4", p.UploadID))
	cmd360p := exec.Command(
		"ffmpeg",
		"-i", p.Path,
		"-vf", "scale=-2:360",
		"-c:v", "libx264",
		"-preset", "fast",
		"-crf", "23",
		"-c:a", "aac",
		"-b:a", "128k",
		"-y",
		output360p,
	)
	cmd360p.Stdout = os.Stdout
	cmd360p.Stderr = os.Stderr
	cmd360p.Start()
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
	output480p := path.Join(fmt.Sprintf("./internal/temp/%s-480p.mp4", p.UploadID))
	cmd480p := exec.Command(
		"ffmpeg",
		"-i", p.Path,
		"-vf", "scale=-2:480",
		"-c:v", "libx264",
		"-preset", "fast",
		"-crf", "23",
		"-c:a", "aac",
		"-b:a", "128k",
		"-y",
		output480p,
	)
	cmd480p.Stdout = os.Stdout
	cmd480p.Stderr = os.Stderr
	cmd480p.Start()
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
	output720p := path.Join(fmt.Sprintf("./internal/temp/%s-720p.mp4", p.UploadID))
	cmd720p := exec.Command(
		"ffmpeg",
		"-i", p.Path,
		"-vf", "scale=-2:720",
		"-c:v", "libx264",
		"-preset", "fast",
		"-crf", "23",
		"-c:a", "aac",
		"-b:a", "128k",
		"-y",
		output720p,
	)
	cmd720p.Stdout = os.Stdout
	cmd720p.Stderr = os.Stderr
	cmd720p.Start()
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
	output1080p := path.Join(fmt.Sprintf("./internal/temp/%s-1080p.mp4", p.UploadID))
	cmd1080p := exec.Command(
		"ffmpeg",
		"-i", p.Path,
		"-vf", "scale=-2:1080",
		"-c:v", "libx264",
		"-preset", "fast",
		"-crf", "23",
		"-c:a", "aac",
		"-b:a", "128k",
		"-y",
		output1080p,
	)
	cmd1080p.Stdout = os.Stdout
	cmd1080p.Stderr = os.Stderr
	cmd1080p.Start()
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
	output1440p := path.Join(fmt.Sprintf("./internal/temp/%s-1440p.mp4", p.UploadID))
	cmd1440p := exec.Command(
		"ffmpeg",
		"-i", p.Path,
		"-vf", "scale=-2:1440",
		"-c:v", "libx264",
		"-preset", "fast",
		"-crf", "23",
		"-c:a", "aac",
		"-b:a", "128k",
		"-y",
		output1440p,
	)
	cmd1440p.Stdout = os.Stdout
	cmd1440p.Stderr = os.Stderr
	cmd1440p.Start()
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
	output4k := path.Join(fmt.Sprintf("./internal/temp/%s-2160p.mp4", p.UploadID))
	cmd4k := exec.Command(
		"ffmpeg",
		"-i", p.Path,
		"-vf", "scale=-2:2160",
		"-c:v", "libx264",
		"-preset", "fast",
		"-crf", "23",
		"-c:a", "aac",
		"-b:a", "128k",
		"-y",
		output4k,
	)
	cmd4k.Stdout = os.Stdout
	cmd4k.Stderr = os.Stderr
	cmd4k.Start()
	return nil
}
