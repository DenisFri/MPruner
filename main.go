package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Path      string `json:"path"`
	DeleteAll bool   `json:"delete_all"`
}

func loadConfig(filename string) (*Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func expandPath(path string) string {
	return os.ExpandEnv(path)
}

func getLastModifiedFile(files []os.DirEntry) (os.DirEntry, error) {
	var lastModifiedFile os.DirEntry
	var lastModifiedTime time.Time

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, err
		}

		if lastModifiedFile == nil || info.ModTime().After(lastModifiedTime) {
			lastModifiedFile = file
			lastModifiedTime = info.ModTime()
		}
	}
	return lastModifiedFile, nil
}

func main() {
	logFile, err := os.OpenFile("cleanup.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatalf("Error closing log file: %v", err)
		}
	}(logFile)
	logger := log.New(logFile, "", log.LstdFlags)

	// Load configuration
	config, err := loadConfig("config.json")
	if err != nil {
		logger.Fatalf("Error loading configuration: %v", err)
	}

	// Expand environment variables in the path
	expandedPath := expandPath(config.Path)
	logger.Printf("Expanded path: %s", expandedPath) // Log the expanded path

	// Verify if the path exists
	if _, err := os.Stat(expandedPath); os.IsNotExist(err) {
		logger.Fatalf("The specified path does not exist: %s", expandedPath)
	}

	// Get list of files in the directory
	files, err := os.ReadDir(expandedPath)
	if err != nil {
		logger.Fatalf("Error reading directory: %v", err)
	}

	if len(files) == 0 {
		logger.Println("No files found in the directory.")
		return
	}

	if config.DeleteAll {
		// Delete all files
		for _, file := range files {
			filePath := filepath.Join(expandedPath, file.Name())
			err := os.Remove(filePath)
			if err != nil {
				logger.Printf("Error deleting file %s: %v", filePath, err)
			} else {
				logger.Printf("Successfully deleted file: %s", filePath)
			}
		}
	} else {
		// Delete only the last modified file
		lastModifiedFile, err := getLastModifiedFile(files)
		if err != nil {
			logger.Fatalf("Error finding the last modified file: %v", err)
		}

		filePath := filepath.Join(expandedPath, lastModifiedFile.Name())
		err = os.Remove(filePath)
		if err != nil {
			logger.Printf("Error deleting file %s: %v", filePath, err)
		} else {
			logger.Printf("Successfully deleted last modified file: %s", filePath)
		}
	}

	logger.Println("Operation completed successfully.")
}
