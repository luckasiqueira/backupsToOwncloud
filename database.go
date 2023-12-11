package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
	"s3backups/models"
)

func DbDump(Data *models.Model, Path *string, LogDir *string) bool {
	fmt.Print("\nDumping database")
	// Loading weber.env file
	err := godotenv.Load()
	if err != nil {
		return false
	}

	// Creating .sql file
	filePath, err := os.Create(os.Getenv("HOME_DIR") + Data.NewPath + ".sql")
	if err != nil {
		fmt.Printf("Error when creating .sql file: %v\n", err)
		os.Exit(1)
	}

	// Run mysqldump command
	rundump := exec.Command("mysqldump", "-u", os.Getenv("DB_USER"), "-p"+os.Getenv("DB_PASS"), "-h", os.Getenv("DB_HOST"), "-P", os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	output, err := rundump.Output()
	if err != nil {
		fmt.Printf("Error when running mysqldump: %v\n", err)
		os.Exit(1)
	}

	// Set filePath variable as output for mysqldump command
	rundump.Stdout = filePath

	_, err = filePath.Write(output)
	if err != nil {
		fmt.Printf("Error when writing to .sql file: %v\n", err)
		os.Exit(1)
	}

	defer filePath.Close()

	return true
}

func DbSync(Data *models.Model, Path *string, LogDir *string, done chan bool) {
	fmt.Print("\nSync database")
	godotenv.Load()

	// Set SQL file's name as same as in DbDump function
	sqlFile := os.Getenv("HOME_DIR") + Data.NewPath + ".sql"

	// Running rclone sync comnand with logging
	dbsync := exec.Command("rclone", "sync", "-P", "--log-level=ERROR", *LogDir, sqlFile, *Path)
	output, err := dbsync.CombinedOutput()

	if err != nil {
		fmt.Print("Error when trying to sync DB to S3.")
		fmt.Print(output)
		return
	}

	// Deletes .sql file after sync
	remove := os.Remove(sqlFile)
	if remove != nil {
		fmt.Print("Error when removing .sql file.")
		return
	}

	done <- true
}
