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

func Init(dirMap map[string][]string, path string) (Navbar, error) {
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
