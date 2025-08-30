package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

func Finalize(uploadId string) string {
	redisInstance := GetRedisInstance()
	ctx := context.Background()
	parts, err := redisInstance.ZRange(ctx, uploadId, 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	file_recreated, _ := os.OpenFile(fmt.Sprintf("./internal/temp/upload-%s-final.mp4", uploadId), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	for _, path := range parts {
		part, _ := os.Open(path)
		io.Copy(file_recreated, part)
		os.Remove(path)
	}
	return fmt.Sprintf("./internal/temp/upload-%s-final.mp4", uploadId)
}
