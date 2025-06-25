package util

import (
	"bytes"
	"text/template"
)

// CompileTemplate compiles the provided template.
func CompileTemplate(text string) *template.Template {
	return template.Must(template.New("").Parse(text))
}

// RenderTemplate renders the provided template to a string.
func RenderTemplate(t *template.Template, data any) string {
	b := &bytes.Buffer{}
	if err := t.Execute(b, data); err != nil {
		panic(err)
	}
	return b.String()
}
