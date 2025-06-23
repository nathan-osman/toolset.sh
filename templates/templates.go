package templates

import (
	"embed"

	"github.com/flosch/pongo2/v6"
	loader "github.com/nathan-osman/pongo2-embed-loader"
)

// C represents context to be rendered with the template.
type C pongo2.Context

var (
	//go:embed templates
	tmplFS embed.FS

	tmplSet = pongo2.NewSet("", &loader.Loader{
		Content: tmplFS,
	})

	// NoContext represents an empty context.
	NoContext = C{}
)

// Render the specified template by name with the provided context.
func Render(name string, ctx C) string {
	t, err := tmplSet.FromFile(name)
	if err != nil {
		panic(err)
	}
	v, err := t.Execute(pongo2.Context(ctx))
	if err != nil {
		panic(err)
	}
	return v
}
