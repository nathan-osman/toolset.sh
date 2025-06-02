package manager

import "fmt"

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
	Name       string   `json:"name"`
	Desc       string   `json:"desc"`
	Params     []*Param `json:"params"`
	RouteNames []string `json:"route_names"`
}

// Input represents data available to the tool.
type Input struct {
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
	tools map[string]Tool
}

// New creates a new tool manager.
func New() *Manager {
	return &Manager{
		tools: make(map[string]Tool),
	}
}

// Register adds the specified tool to the map.
func (m *Manager) Register(t Tool) {
	for _, k := range t.Meta().RouteNames {
		m.tools[k] = t
	}
}

// Get returns the specified tool based on its name.
func (m *Manager) Get(name string) (Tool, error) {
	t, ok := m.tools[name]
	if !ok {
		return nil, fmt.Errorf("\"%s\" is not a recognized tool", name)
	}
	return t, nil
}
