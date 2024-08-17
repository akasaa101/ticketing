package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestConfig_SuccessfulLoad(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")

	value := Config("TEST_KEY")
	expected := "test_value"

	if value != expected {
		t.Errorf("Expected %s, got %s", expected, value)
	}
}

func TestConfig_MissingKey(t *testing.T) {
	value := Config("MISSING_KEY")
	expected := ""

	if value != expected {
		t.Errorf("Expected empty string, got %s", value)
	}
}

func TestConfig_LoadError(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	rootDir := filepath.Dir(filepath.Dir(wd))
	envPath := filepath.Join(rootDir, ".env")
	envBackupPath := filepath.Join(rootDir, ".env.bak")

	fmt.Printf("Original .env path: %s\n", envPath)
	fmt.Printf("Backup .env path: %s\n", envBackupPath)

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		file, err := os.Create(envPath)
		if err != nil {
			t.Fatalf("Failed to create temporary .env file: %v", err)
		}
		file.Close()
		defer os.Remove(envPath)
	}

	err = os.Rename(envPath, envBackupPath)
	if err != nil {
		t.Fatalf("Failed to rename .env: %v", err)
	}
	defer os.Rename(envBackupPath, envPath)

	Config("ANY_KEY")
}
