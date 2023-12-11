package main

import (
	"github.com/joho/godotenv"
	"os"
	"os/exec"
	"s3backups/models"
)

var Data = &models.Model{}

func main() {
	exec.Command("pwd")
	// Define some dates
	models.Now(Data)
	models.Past(Data)

	// Set our new directory into S3 bucket
	godotenv.Load()
	var p = os.Getenv("S3_PATH") + Data.NewPath
	var Path = &p

	// Define our log file name
	var l = "--log-file=" + os.Getenv("HOME_DIR") + "log-" + Data.NewPath + ".txt"
	var LogDir = &l

	waitDb := make(chan bool)
	waitFiles := make(chan bool)

	// Runs DB Dump and Sync
	DbDump(Data, Path, LogDir)

	// Runs DB and File's sync at same time using go routines to ensure they both will run
	go DbSync(Data, Path, LogDir, waitDb)
	go FilesSync(Data, Path, LogDir, waitFiles)

	// Hold executions, waiting for Db and File sync's completion
	<-waitDb
	<-waitFiles

	// Remove older backup directories
	defer Rotate(Data)
}
