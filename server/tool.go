package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/toolset.sh/manager"
)

func (s *Server) tool(c *gin.Context) {

	// Determine the name of the tool to run
	name := c.Param("name")

	// Attempt to find the tool
	t, err := s.manager.Get(name)
	if err != nil {
		s.sendError(c, err.Error())
		return
	}

	// TODO: parameters

	// The final statement in this function body executes the tool - it may
	// panic, so handle the actual output in the defer below

	var o manager.Output

	defer func() {
		if r := recover(); r != nil {
			var msg string
			switch v := r.(type) {
			case error:
				msg = v.Error()
			case string:
				msg = v
			default:
				msg = "an unknown internal error has occurred"
			}
			s.sendError(c, msg)
		} else {
			s.sendOutput(c, o)
		}
	}()

	o = t.Run(&manager.Input{})
}
