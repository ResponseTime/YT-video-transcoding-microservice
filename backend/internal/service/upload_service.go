package service

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func UploadService(uploadId string, chunk multipart.File, part string, end string) {
	var file *os.File
	// partn, _ := strconv.Atoi(part)
	file, err := os.OpenFile(fmt.Sprintf("./internal/temp/%s.mp4", uploadId), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := io.Copy(file, chunk); err != nil {
		log.Fatal(err)
	}
	if end == "True" {
		fmt.Println("✅ All chunks uploaded, ready to merge")
	}
}
