package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/nathan-osman/toolset.sh/templates"
)

type outputFormat int

const (
	outputJson outputFormat = iota
	outputHtml
	outputText
)

func (s *Server) sendOutput(c *gin.Context, r manager.Output) {

	acceptHeader := c.GetHeader("Accept")

	// Determine the appropriate output format from the Accept header
	// - if this fails, send a text response since we wouldn't know what type
	//   to send to the client anyway - this function cannot fail or panic &
	//   therefore can be used in the regular panic handler
	t, err := parseAcceptHeader(acceptHeader)
	if err != nil {
		c.AbortWithError(
			http.StatusBadRequest,
			fmt.Errorf("invalid Accept header '%s'", acceptHeader),
		)
		return
	}

	// Check for cURL with the default Accept: header
	if strings.HasPrefix(c.GetHeader("User-Agent"), "curl/") &&
		c.GetHeader("Accept") == "*/*" {
		t = outputText
	}

	var v []byte

	// Set up the response for the selected type
	switch t {
	case outputJson:
		c.JSON(http.StatusOK, r)
		return
	case outputHtml:
		v = []byte(templates.Render("templates/base.html", templates.C{
			"value": r.Html(),
		}))
		c.Header("Content-Type", "text/html; charset=utf-8")
	case outputText:
		v = []byte(r.Text() + "\n")
		c.Header("Content-Type", "text/plain; charset=utf-8")
	}

	// Write the response
	c.Header("Content-Length", strconv.Itoa(len(v)))
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(v)
}

type outputError struct {
	Error string `json:"error"`
}

func (o *outputError) Text() string {
	return o.Error
}

func (o *outputError) Html() string {
	return templates.Render("templates/error.html", templates.C{
		"msg": o.Error,
	})
}

// TODO: status code for errors

func (s *Server) sendError(c *gin.Context, msg string) {
	s.sendOutput(c, &outputError{Error: msg})
}
