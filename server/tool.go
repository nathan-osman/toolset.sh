package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/toolset.sh/registry"
	"github.com/nathan-osman/toolset.sh/templates"
)

type toolOutput struct {
	meta   *registry.Meta
	output registry.Output
}

func (t *toolOutput) Text() string {
	return t.output.Text()
}

func (t *toolOutput) Html() string {
	return templates.Render(
		"templates/tool.html",
		templates.C{
			"meta":   t.meta,
			"output": t.output.Html(),
		},
	)
}

func (s *Server) tool(c *gin.Context) {

	// Determine the name of the tool to run
	name := c.Param("name")

	// Attempt to find the tool
	t, n, err := registry.Get(name)
	if err != nil {
		panic(err)
	}

	// If the canonical name was not used, indicate that we want to redirect
	// to it; this is only a suggestion since we don't want to do that for
	// text clients
	if n != "" {
		q := c.Request.URL.RawQuery
		if len(q) != 0 {
			q = "?" + q
		}
		c.Set(contextRedirect, "/"+n+q)
	}

	// Run the tool and obtain its output
	o := t.Run(
		convertContextToInput(c, t),
	)

	// Send the output to the client
	s.sendOutput(
		c,
		&toolOutput{
			meta:   t.Meta(),
			output: o,
		},
	)
}
