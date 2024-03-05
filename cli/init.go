package cli

import (
	"log"
	"os"
	"path/filepath"
)

// Init initialize new project repository in the given path
// with empty "index", "build", "templates" and "static" folders
func Init(path string) error {
	paths := []string{"index", "build", "templates", "static"}
	var dirPath string

	// Make all directories
	for _, p := range paths {
		dirPath = filepath.Join(path, p)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}
	fp, _ := filepath.Abs(path)
	log.Println("blender: created Project at ", fp)
	return nil
}
