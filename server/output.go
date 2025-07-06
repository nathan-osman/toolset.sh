package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/toolset.sh/registry"
	"github.com/nathan-osman/toolset.sh/templates"
)

type outputFormat int

const (
	contextStatusCode = "status_code"
	contextRedirect   = "redirect"

	outputJson outputFormat = iota
	outputHtml
	outputText
)

func (s *Server) sendOutput(c *gin.Context, r registry.Output) {

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

	// Check for cURL / wget with the default Accept: header
	var (
		a      = c.GetHeader("User-Agent")
		isCurl = strings.HasPrefix(a, "curl/")
		isWget = strings.HasPrefix(a, "Wget/")
	)
	if (isCurl || isWget) && c.GetHeader("Accept") == "*/*" {
		t = outputText
	}

	// Allow the output= param to override everything
	switch c.Query("output") {
	case "json":
		t = outputJson
	case "html":
		t = outputHtml
	case "text":
		t = outputText
	}

	// If the output format is HTML and a redirect was set, use it
	if t == outputHtml {
		if v, ok := c.Get(contextRedirect); ok {
			c.Redirect(
				http.StatusMovedPermanently,
				v.(string),
			)
			return
		}
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

	// Get the status code
	statusCode := http.StatusOK
	if v, ok := c.Get(contextStatusCode); ok {
		statusCode = v.(int)
	}

	// Write the response
	c.Header("Content-Length", strconv.Itoa(len(v)))
	c.Writer.WriteHeader(statusCode)
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

func (s *Server) sendError(c *gin.Context, msg string) {
	c.Set(contextStatusCode, http.StatusInternalServerError)
	s.sendOutput(c, &outputError{Error: msg})
}
