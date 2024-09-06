package utils

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// CleanDirectory deletes files from a directory based on the deleteAll flag
func CleanDirectory(dirPath string, deleteAll bool, logger *log.Logger) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		logger.Println("No files found in the directory.")
		return nil
	}

	if deleteAll {
		for _, file := range files {
			filePath := filepath.Join(dirPath, file.Name())
			err := os.Remove(filePath)
			if err != nil {
				logger.Printf("Error deleting file %s: %v", filePath, err)
			} else {
				logger.Printf("Successfully deleted file: %s", filePath)
			}
		}
	} else {
		// Delete only the last modified file
		lastModifiedFile, err := getLastModifiedFile(files, dirPath)
		if err != nil {
			return err
		}
		filePath := filepath.Join(dirPath, lastModifiedFile.Name())
		err = os.Remove(filePath)
		if err != nil {
			logger.Printf("Error deleting file %s: %v", filePath, err)
		} else {
			logger.Printf("Successfully deleted last modified file: %s", filePath)
		}
	}

	return nil
}

// getLastModifiedFile returns the last modified file in the directory
func getLastModifiedFile(files []os.DirEntry, dirPath string) (os.DirEntry, error) {
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
