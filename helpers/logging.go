package helpers

import (
	"log"
	"os"
	"path/filepath"
)

func Log(message string, logFile string) {
	// Find the project root by searching for go.mod
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting working directory: " + err.Error())
	}

	rootPath := workingDir
	for {
		if _, err := os.Stat(filepath.Join(rootPath, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(rootPath)
		if parent == rootPath {
			// Reached the root of the file system
			rootPath = workingDir
			break
		}
		rootPath = parent
	}

	// Join the root path with the relative log file path
	absPath := filepath.Join(rootPath, logFile)

	// Create directory if it doesn't exist
	dir := filepath.Dir(absPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal("error creating log directory: " + err.Error())
		}
	}

	file, err := os.OpenFile(absPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("error openning the file: " + err.Error())
	}

	log.SetOutput(file)

	log.Println(message)
}
