package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) tool(c *gin.Context) {

	// Determine the name of the tool to run
	name := c.Param("name")

	// Attempt to find the tool
	t, err := s.manager.Get(name)
	if err != nil {
		panic(err)
	}

	// Send the output to the client
	s.sendOutput(
		c,
		t.Run(
			convertContextToInput(c, t),
		),
	)
}
