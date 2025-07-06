package registry

import (
	"github.com/gin-gonic/gin"
)

// Input represents data available to the tool.
type Input struct {
	C      *gin.Context
	Params map[string]string
}

// Output is used to format the results from a tool.
type Output interface {
	Text() string
	Html() string
}
