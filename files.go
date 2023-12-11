package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
	"s3backups/models"
)

func FilesSync(Data *models.Model, Path *string, LogDir *string, done chan bool) {
	fmt.Print("\nFiles")
	godotenv.Load()

	// Executes rclone sync
	sync := exec.Command("rclone", "sync", "-P", "--log-level=ERROR", *LogDir, os.Getenv("FILES_DIR"), *Path)
	output, err := sync.CombinedOutput()

	if err != nil {
		fmt.Print(string(output))
		return
	}
	done <- true
}
