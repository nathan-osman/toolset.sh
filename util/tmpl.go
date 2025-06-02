package util

import (
	"bytes"
	"html/template"
)

// MustCreateTemplate is a shortcut for
// template.Must(template.New().Parse()).
func MustCreateTemplate(tmpl string) *template.Template {
	return template.Must(template.New("").Parse(tmpl))
}

// MustRenderTemplateToString renders a template to string and panics if it
// fails.
func MustRenderTemplateToString(t *template.Template, data any) string {
	b := bytes.Buffer{}
	if err := t.Execute(&b, data); err != nil {
		panic(err)
	}
	return b.String()
}
