package cli

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/awalvie/go-blender/logging"
)

// Build generate default directory structure in the given path
func Build(buildPath string) error {
	logging.InfoLogger.Println("Generating directory tree")

	fileMap, err := generateFileMap(buildPath)
	if err != nil {
		logging.ErrorLogger.Println("Failed to generate file map for given directory")
		return err
	}

	for k, v := range fileMap {
		logging.InfoLogger.Println("\nKey: ", k, "\nValues: ", v)
	}
}

func generateFileMap(root string) (map[string][]string, error) {
	fileMap := map[string][]string{}
	var indexDir string

	if _, err := os.Stat(root + "/index"); err != nil {
		logging.ErrorLogger.Println("Index directory doesn't exist in path")
		return nil, err
	} else {
		indexDir = root + "/index"
	}

	err := filepath.Walk(indexDir, func(path string, fi os.FileInfo, err error) error {
		fileNames := []string{}

		if fi.IsDir() == true {
			files, err := ioutil.ReadDir(path)
			if err != nil {
				return err
			}
			for _, f := range files {
				fileNames = append(fileNames, f.Name())
			}
			fileMap[path] = fileNames

		} else if fi.Name() != "_index.md" {
			fileMap[path] = nil
		}
		return nil
	})
	return fileMap, err
}
