package cli

import (
	"os"
	"path/filepath"

	"github.com/awalvie/go-blender/logging"
)

// Init initialize new project repository in the given path
// with empty "index", "build", "templates" and "static" folders
func Init(path string) error {
	paths := []string{"index", "build", "templates", "static"}
	var dirPath string

	for _, p := range paths {
		dirPath = filepath.Join(path, p)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}
	fp, _ := filepath.Abs(path)
	logging.InfoLogger.Println("blender: created Project at ", fp)
	return nil
}
