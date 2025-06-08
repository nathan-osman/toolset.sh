package server

import (
	"bytes"
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed templates
	tmplFS embed.FS
)

func render(tmpl string, v any) string {
	t, err := template.ParseFS(tmplFS, tmpl)
	if err != nil {
		panic(err)
	}
	b := &bytes.Buffer{}
	if err := t.Execute(b, v); err != nil {
		panic(err)
	}
	return b.String()
}

func renderIndex(v string) string {
	return render("templates/base.html", template.HTML(v))
}

type indexOutput struct{}

func (i *indexOutput) Text() string {
	return "Coming soon!"
}

func (i *indexOutput) Html() string {
	return renderIndex(render("templates/index.html", nil))
}

func (s *Server) index(c *gin.Context) {
	s.sendOutput(c, &indexOutput{})
}
