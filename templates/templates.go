package templates

import (
	"embed"

	"github.com/flosch/pongo2/v6"
	loader "github.com/nathan-osman/pongo2-embed-loader"
)

var (
	//go:embed templates
	tmplFS embed.FS

	tmplSet = pongo2.NewSet("", &loader.Loader{
		Content: tmplFS,
	})

	// NoContext provides a convenient empty context block.
	NoContext = pongo2.Context{}
)

// Render the specified template by name with the provided context.
func Render(name string, ctx pongo2.Context) string {
	t, err := tmplSet.FromFile(name)
	if err != nil {
		panic(err)
	}
	v, err := t.Execute(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
