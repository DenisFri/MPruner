package main

import (
	"log"
	"os"

	"MPruner/config"
	"MPruner/utils"
)

func main() {
	logger, err := utils.InitLogger("cleanup.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Load configuration from config.json
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		logger.Fatalf("Error loading configuration: %v", err)
	}

	// Map to track processed directories
	processedDirs := make(map[string]bool)

	for _, dirCfg := range cfg.Directories {
		expandedPath := os.ExpandEnv(dirCfg.Path)
		logger.Printf("Expanded path: %s", expandedPath)

		if processedDirs[expandedPath] {
			logger.Printf("Skipping already processed directory: %s", expandedPath)
			continue
		}

		processedDirs[expandedPath] = true

		// Check if the path exists
		if _, err := os.Stat(expandedPath); os.IsNotExist(err) {
			logger.Printf("The specified path does not exist: %s", expandedPath)
			continue
		}

		// Perform the cleanup (delete files) in the directory
		err = utils.CleanDirectory(expandedPath, dirCfg.DeleteAll, logger)
		if err != nil {
			logger.Printf("Error cleaning directory %s: %v", expandedPath, err)
		} else {
			logger.Printf("Successfully cleaned directory: %s", expandedPath)
		}
	}

	logger.Println("Operation completed successfully.")
}
