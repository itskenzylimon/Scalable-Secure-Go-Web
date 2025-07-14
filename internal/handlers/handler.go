package handlers

import (
	"github.com/go-playground/validator/v10"
	"log"
	"os"
	"path/filepath"
)

var validateProduct = validator.New()

var validateCategory = validator.New()

var validateBrand = validator.New()

func SetupLogFile() *os.File {
	logDir := "logs"
	logFile := "server.log"

	// Create logs/ directory if not exists
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("❌ Failed to create log directory: %v", err)
	}

	// Open or create the log file
	file, err := os.OpenFile(filepath.Join(logDir, logFile), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("❌ Failed to open log file: %v", err)
	}

	return file
}
