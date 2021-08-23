package cli

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/awalvie/go-blender/markdown"
	"github.com/awalvie/go-blender/utils"
)

const (
	BUILD_DIR     = "build/"
	INDEX_DIR     = "index/"
	TEMPLATES_DIR = "templates/"
	STATIC_DIR    = "static/"
	EXT_MD        = ".md"
	_INDEX        = "_index.md"
	PATH_DELIM    = "/"
)

type Navbar struct {
	Parent  []Element
	Current []Element
	Child   []Element
}

type Element struct {
	Href string
	Name string
}

func genNavbar(dirMap map[string][]string, path string) (Navbar, error) {
	// TODO: Handle top level index
	navbar := Navbar{}

	// get all child nodes from map
	childNodes := dirMap[path]
	for _, v := range childNodes {
		if v != "_index.md" {
			name := strings.TrimSuffix(v, ".md")
			href := utils.ToHTML(v)
			element := Element{href, name}
			navbar.Child = append(navbar.Child, element)
		}
	}

	// get all current nodes from map
	currentDir := filepath.Dir(path)
	currentNodes := dirMap[currentDir]

	for _, v := range currentNodes {
		if v != "_index.md" {
			var element Element
			if base := filepath.Base(path); base == v {
				name := strings.TrimSuffix(v, ".md") + "/"
				href := utils.ToHTML(v)
				element = Element{href, name}
				navbar.Current = append(navbar.Current, element)
			} else {
				name := strings.TrimSuffix(v, ".md")
				href := utils.ToHTML(v)
				element = Element{href, name}
				navbar.Current = append(navbar.Current, element)
			}

		}
	}

	// get all parent nodes from map
	parentDir := filepath.Dir(currentDir)
	parentNodes := dirMap[parentDir]

	for _, v := range parentNodes {
		if v != "_index.md" {
			var element Element

			if base := filepath.Base(currentDir); base == v {
				name := strings.TrimSuffix(v, ".md") + "/"
				href := utils.ToHTML(v)
				element = Element{href, name}
				navbar.Parent = append(navbar.Parent, element)
			} else {
				name := strings.TrimSuffix(v, ".md")
				href := utils.ToHTML(v)
				element = Element{href, name}
				navbar.Parent = append(navbar.Parent, element)
			}
		}
	}

	return navbar, nil
}

// initDirMap returns a map with:
// key: path to every file/directory inside the given root
// value: first level files/directories if the key is a directory
// Ignores '_index.md' since that signifies the directory page itself
func initDirMap(root string) (map[string][]string, error) {

	fileMap := map[string][]string{}
	indexDir := filepath.Join(
		root,
		PATH_DELIM,
		INDEX_DIR,
	)

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

// renderFiles needs a directory map and buildPath as arguments.
// It parses all .md files in the buildPath/site and renders them into
// respective htmls in the buildPath/build directory
func renderFiles(dirMap map[string][]string, buildPath string) error {

	// if key is directory, parse _index.md and render as directory.html
	// if key is file, parse markdown and reader it as file.html
	for path := range dirMap {

		// if key has '.md' it's a file, directly parse file markdown
		if strings.Contains(path, EXT_MD) {

			mdData, metadata, err := markdown.RenderMD(path)
			if err != nil {
				return err
			}

			// created file paths
			fileName := filepath.Base(path)
			filePath := filepath.Join(
				buildPath,
				BUILD_DIR,
				strings.Replace(fileName, "md", "html", 1),
			)

			// templateDir
			templateDir := filepath.Join(
				buildPath,
				TEMPLATES_DIR,
			)

			// generate Navbar for website
			navbar, err := genNavbar(dirMap, path)
			if err != nil {
				return err
			}

			// render HTML
			data := struct {
				Body string
				Meta map[string]interface{}
				Nav  Navbar
			}{mdData.String(), metadata, navbar}

			err = markdown.RenderHTML(filePath, templateDir, metadata, data)
			if err != nil {
				return err
			}

		} else {
			// check if the directory has a a _index.md file to represent
			// the folder page and parse it
			indexFile := filepath.Join(
				path,
				PATH_DELIM,
				_INDEX,
			)

			if _, err := utils.Exists(indexFile); err != nil {
				return err
			}
		}
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
	buildDir := filepath.Join(
		buildPath,
		PATH_DELIM,
		BUILD_DIR,
	)

	if err := utils.Clean(buildDir); err != nil {
		return err
	}

	// Get the file map
	fileMap, err := initDirMap(buildPath)
	if err != nil {
		return err
	}

	// Parse files/folders in map and renders them in HTML
	if err := renderFiles(fileMap, buildPath); err != nil {
		return err
	}

	// Copy the static directory into build
	// ex: build/static/
	staticDir := filepath.Join(
		buildPath,
		PATH_DELIM,
		STATIC_DIR,
	)

	buildStatic := filepath.Join(buildDir, STATIC_DIR)

	os.Mkdir(buildStatic, 0755)
	if err := utils.CopyDir(staticDir, buildStatic); err != nil {
		return err
	}

	return nil
}
