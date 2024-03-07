package navbar

import (
	"path/filepath"
	"strings"

	"github.com/awalvie/go-blender/utils"
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

func addElements(nodes []string, target *[]Element, basePath string) {
	for _, v := range nodes {
		// Ignore the directory page of the parent directory's parent
		if v == "_index.md" {
			continue
		}

		name := strings.ReplaceAll(strings.TrimSuffix(v, ".md"), "_", " ")
		href := utils.ToHTML(v)

		if base := filepath.Base(basePath); base == v {
			name += "/"
		}

		*target = append(*target, Element{Href: href, Name: name})
	}
}

func (nav *Navbar) Init(dirMap map[string][]string, path string) error {
	childNodes := dirMap[path]
	addElements(childNodes, &nav.Child, path)

	currentDir := filepath.Dir(path)
	currentNodes := dirMap[currentDir]
	addElements(currentNodes, &nav.Current, path)

	parentDir := filepath.Dir(currentDir)
	parentNodes := dirMap[parentDir]
	addElements(parentNodes, &nav.Parent, currentDir)

	return nil
}
