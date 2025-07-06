package registry

import (
	"fmt"
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
	Category       string   `json:"category"`
	Name           string   `json:"name"`
	Desc           string   `json:"desc"`
	Params         []*Param `json:"params"`
	RouteName      string   `json:"route_name"`
	AlternateNames []string `json:"alternate_names"`
}

// Tool is the interface that all tools must implement.
type Tool interface {
	Meta() *Meta
	Run(*Input) Output
}

var (

	// ToolMap contains a map of all tools.
	ToolMap = map[string]Tool{}

	// AlternateNameMap maps alternate names to tool names.
	AlternateNameMap = map[string]string{}
)

// Register registers the specified tool. Note that this should only be called
// on initialization to maintain thread safety.
func Register(t Tool) {
	var (
		m = t.Meta()
		n = m.RouteName
	)
	ToolMap[n] = t
	for _, k := range m.AlternateNames {
		AlternateNameMap[k] = n
	}
	for _, c := range Categories {
		if c.Name == m.Category {
			c.Tools = append(c.Tools, t)
			break
		}
	}
}

// Get attempts to retrieve the specified tool. If an alternate name is used,
// the original (canonical) name is also returned.
func Get(name string) (Tool, string, error) {
	t, ok := ToolMap[name]
	if !ok {
		n, ok := AlternateNameMap[name]
		if !ok {
			return nil, "", fmt.Errorf("\"%s\" is not a recognized tool", name)
		}
		return ToolMap[n], n, nil
	}
	return t, "", nil
}
