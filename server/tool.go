package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	// Send the output to the client
	s.sendOutput(
		c,
		t.Run(
			convertContextToInput(c, t),
		),
	)
}
