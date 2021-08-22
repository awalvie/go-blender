package utils

import (
	"os"
	"path"
	"path/filepath"
)

// Cleans a given directory, removing all files and subdirs.
func Clean(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// Exists reports whether the named file or directory exists.
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	return true, nil
}

// Modify any string to have a .html extension
func ToHTML(name string) string {
	ext := path.Ext(name)
	htmlName := name[0:len(name)-len(ext)] + ".html"

	return htmlName
}
