package cli

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/awalvie/go-blender/logging"
	"github.com/awalvie/go-blender/utils"
)

const (
	BUILD     = "/build"
	INDEX     = "/index"
	TEMPLATES = "/templates"
	STATIC    = "/static"
	EXT_MD    = ".md"
	_INDEX    = "_index.md"
)

// initDirMap returns a map with:
// key: path to every file/directory inside the given root
// value: first level files/directories if the key is a directory
// Ignores '_index.md' since that signifies the directory page itself
func initDirMap(root string) (map[string][]string, error) {

	fileMap := map[string][]string{}
	indexDir := root + INDEX

	if _, err := utils.Exists(indexDir); err != nil {
		return nil, err
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

		} else if fi.Name() != _INDEX {
			fileMap[path] = nil
		}
		return nil
	})
	return fileMap, err
}

func renderFiles(fileMap map[string][]string) error {
	logging.InfoLogger.Println("Rendering Files")

	for k, v := range fileMap {
		logging.InfoLogger.Println("Key:", k, "Value:", v)
	}
	return nil
}

// Build does the following:
// * Generates directory map from the index directory
// * Parses files/directories in the dir map
// * Generates static content
// * Copies it into the build folder
func Build(buildPath string) error {

	// Clean the build directory.
	buildDir := buildPath + BUILD
	logging.InfoLogger.Println("Cleaning 'build' directory")
	if err := utils.Clean(buildDir); err != nil {
		return err
	}

	// Get the file map
	fileMap, err := initDirMap(buildPath)
	if err != nil {
		return err
	}

	// Parse files/folders in map and renders them in HTML
	if err := renderFiles(fileMap); err != nil {
		return err
	}

	return nil
}
