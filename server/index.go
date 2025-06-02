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
