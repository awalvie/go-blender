package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// generate directory structure
func blenderBuild(buildPath string) {
	InfoLogger.Println("Generating directory tree")
	fileMap, err := genFileMap(buildPath)
	if err != nil {
		ErrorLogger.Println(err)
	}
	InfoLogger.Println("Building Site")

	for k, v := range fileMap {
		InfoLogger.Println("\nKey: ", k, "\nValues", v)
	}
}

func genFileMap(root string) (map[string][]string, error) {
	fileMap := map[string][]string{}
	var indexDir string

	if _, err := os.Stat(root + "/index"); err != nil {
		ErrorLogger.Println("Index directory doesn't exist in path")
	} else {
		indexDir = root + "/index"
	}

	err := filepath.Walk(indexDir, func(path string, fi os.FileInfo, err error) error {
		fileNames := []string{}

		if fi.IsDir() == true {
			files, err := ioutil.ReadDir(path)
			if err != nil {
				ErrorLogger.Println(err)
				return nil
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
