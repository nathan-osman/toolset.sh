package manager

import "fmt"

// Output is used to format the results from a tool.
type Output interface {
	Text() string
	Html() string
}

// Tool is the interface that all tools must implement.
type Tool interface {
	Name() string
	Desc() string
	RouteNames() []string
	Run() (Output, error)
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
	for _, k := range t.RouteNames() {
		m.tools[k] = t
	}
}

// Run executes the specified tool.
func (m *Manager) Run(name string) (Output, error) {
	t, ok := m.tools[name]
	if !ok {
		return nil, fmt.Errorf("\"%s\" is not a recognized tool")
	}
	return t.Run()
}
