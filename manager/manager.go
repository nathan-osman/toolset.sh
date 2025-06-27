package manager

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Option represents a valid value for a parameter.
type Option struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

// Param provides parameter information.
type Param struct {
	Name    string    `json:"name"`
	Desc    string    `json:"desc"`
	Default string    `json:"default"`
	Options []*Option `json:"options"`
}

// Meta provides information about a tool.
type Meta struct {
	Name           string   `json:"name"`
	Desc           string   `json:"desc"`
	Params         []*Param `json:"params"`
	RouteName      string   `json:"route_name"`
	AlternateNames []string `json:"alternate_names"`
}

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

// Tool is the interface that all tools must implement.
type Tool interface {
	Meta() *Meta
	Run(*Input) Output
}

// Manager keeps track of the available tools and provides access by command name.
type Manager struct {
	tools          map[string]Tool
	alternateNames map[string]string
}

// New creates a new tool manager.
func New() *Manager {
	return &Manager{
		tools:          make(map[string]Tool),
		alternateNames: make(map[string]string),
	}
}

// Register adds the specified tool to the map.
func (m *Manager) Register(t Tool) {
	n := t.Meta().RouteName
	m.tools[n] = t
	for _, k := range t.Meta().AlternateNames {
		m.alternateNames[k] = n
	}
}

// Get returns the specified tool based on its name. If an alternate name for
// the tool was used, the default is returned - perfect for a 301 redirect.
func (m *Manager) Get(name string) (Tool, string, error) {
	t, ok := m.tools[name]
	if !ok {
		n, ok := m.alternateNames[name]
		if !ok {
			return nil, "", fmt.Errorf("\"%s\" is not a recognized tool", name)
		}
		return m.tools[n], n, nil
	}
	return t, "", nil
}
