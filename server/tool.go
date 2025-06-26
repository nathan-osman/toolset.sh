package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/nathan-osman/toolset.sh/templates"
)

type toolOutput struct {
	meta   *manager.Meta
	output manager.Output
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
	t, n, err := s.manager.Get(name)
	if err != nil {
		panic(err)
	}

	// If the canonical name was not used, redirect to it
	if t == nil {
		q := c.Request.URL.RawQuery
		if len(q) != 0 {
			q = "?" + q
		}
		c.Redirect(
			http.StatusMovedPermanently,
			"/"+n+q,
		)
		return
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
