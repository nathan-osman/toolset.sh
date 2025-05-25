package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) index(c *gin.Context) {

	// TODO: show helpful intro

	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("Coming soon!\n"))
}

func (s *Server) tool(c *gin.Context) {
	name := c.Param("name")
	r, err := s.manager.Run(name)
	if err != nil {
		s.sendError(c, err)
		return
	}
	s.sendOutput(c, r)
}
