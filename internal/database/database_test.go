package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestConnect(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	rootDir := filepath.Dir(filepath.Dir(wd))
	envPath := filepath.Join(rootDir, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	Connect()

	assert.NotNil(t, DB.Db, "Database connection should be established")
}
