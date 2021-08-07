package main

import (
	"os"
	"path/filepath"
)

// blenderInit initialize new project repository in the given path
func blenderInit(path string) {
	paths := []string{"index", "build", "templates", "static"}
	var dirPath string
	for _, p := range paths {
		dirPath = filepath.Join(path, p)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			ErrorLogger.Println(err)
			return
		}
	}
	fp, _ := filepath.Abs(path)
	InfoLogger.Println("Created Project at ", fp)
}
