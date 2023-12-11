package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
	"s3backups/models"
)

func Rotate(Data *models.Model) {
	fmt.Print("\nRotate")
	godotenv.Load()
	// Remove older backups
	for i := range Data.OldPath {
		oldPath := os.Getenv("S3_PATH") + Data.OldPath[i]
		rotate := exec.Command("rclone", "purge", oldPath)
		output, err := rotate.CombinedOutput()

		if err != nil {
			fmt.Print("\nError when removing older backups.")
			fmt.Sprint(output)
		}
	}
}
