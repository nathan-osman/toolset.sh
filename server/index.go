package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/toolset.sh/templates"
)

type indexOutput struct{}

func (i *indexOutput) Text() string {
	return templates.Render("templates/index.txt", templates.NoContext)
}

func (i *indexOutput) Html() string {
	return templates.Render("templates/index.html", templates.NoContext)
}

func (s *Server) index(c *gin.Context) {
	s.sendOutput(c, &indexOutput{})
}
