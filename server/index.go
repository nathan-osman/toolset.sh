package server

import (
	"github.com/gin-gonic/gin"
)

type indexOutput struct{}

func (i *indexOutput) Text() string {
	return "Coming soon!"
}

func (i *indexOutput) Html() string {
	return `<div class="text-2xl">Coming soon!</div>`
}

func (s *Server) index(c *gin.Context) {
	s.sendOutput(c, &indexOutput{})
}
