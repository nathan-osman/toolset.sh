package manager

// Tool is the interface that all tools must implement.
type Tool interface {
	Name() string
	Desc() string
	Commands() []string
	Run() (string, string, error)
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
	for _, k := range t.Commands() {
		m.tools[k] = t
	}
}
