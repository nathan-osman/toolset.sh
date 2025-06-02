package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/toolset.sh/manager"
)

type outputType int

const (
	outputJson outputType = iota
	outputHtml
	outputText
)

func (s *Server) sendOutput(c *gin.Context, r manager.Output) {

	var t outputType = outputHtml

	// TODO: check the accept header

	// Check for cURL with the default Accept: header
	if strings.HasPrefix(c.GetHeader("User-Agent"), "curl/") &&
		c.GetHeader("Accept") == "*/*" {
		t = outputText
	}

	// Write the output for the final decision
	switch t {
	case outputJson:
		c.JSON(http.StatusOK, r)
	case outputHtml:
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(r.Html()))
	case outputText:
		c.Header("Content-Type", "text/plain; charset=utf-8")
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(r.Text()))
		c.Writer.Write([]byte("\n"))
	}
}

type outputError struct {
	Error string `json:"error"`
}

func (o *outputError) Text() string {
	return o.Error
}

func (o *outputError) Html() string {
	return o.Text()
}

// TODO: status code for errors

func (s *Server) sendError(c *gin.Context, msg string) {
	s.sendOutput(c, &outputError{Error: msg})
}
