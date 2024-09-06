package utils

import (
	"log"
	"os"
)

func InitLogger(logFilePath string) (*log.Logger, error) {
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := log.New(logFile, "", log.LstdFlags)

	return logger, nil
}
