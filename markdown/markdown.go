package markdown

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func RenderMD(filepath string) (bytes.Buffer, map[string]interface{}, error) {
	// initialize goldmark
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			meta.Meta,
			extension.Strikethrough,
			extension.GFM,
			extension.Footnote,
		),
	)

	// read contents of the file into a buffer
	mdData, err := os.ReadFile(filepath)
	if err != nil {
		return bytes.Buffer{}, nil, err
	}

	// render markdown
	var htmlData bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert(mdData, &htmlData, parser.WithContext(context)); err != nil {
		return bytes.Buffer{}, nil, err
	}

	metadata := meta.Get(context)

	return htmlData, metadata, nil
}

// Renders out.HTML into dst html file, using the template specified
// in the frontmatter. data is the template struct.
func RenderHTML(destination, templateDir string, meta map[string]interface{}, data interface{}) error {
	metaTemplate := meta["template"]

	if metaTemplate == nil {
		metaTemplate = "default.html"
	}

	t, err := template.New("").ParseGlob(filepath.Join(templateDir, "*.html"))

	if err != nil {
		return err
	}

	file, err := os.Create(destination)
	if err != nil {
		return err
	}

	if err = t.ExecuteTemplate(file, metaTemplate.(string), data); err != nil {
		return err
	}
	return nil
}
