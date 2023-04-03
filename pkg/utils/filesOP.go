package utils

import (
	"log"
	"os"
	"path/filepath"
)

func ReadFileStr(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ProcessPath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return filepath.Dir(ex), nil
}
