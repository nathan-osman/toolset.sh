package registry

// Category stores metadata for categories and their list of tools.
type Category struct {
	Name  string `json:"string"`
	Tools []Tool `json:"tools"`
}

// Categories maintains a list of all recognized categories
var Categories = []*Category{}

func addCategory(name string) string {
	Categories = append(Categories, &Category{Name: name})
	return name
}

var (
	CategoryDateTime    = addCategory("Date / Time")
	CategoryEngineering = addCategory("Engineering")
	CategoryMath        = addCategory("Math")
	CategoryNetwork     = addCategory("network")
	CategoryProgramming = addCategory("programming")
)
