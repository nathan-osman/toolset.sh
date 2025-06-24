package pi

import (
	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/nathan-osman/toolset.sh/templates"
)

var (
	meta = &manager.Meta{
		Name:       "Pi",
		Desc:       "return the value for Pi",
		Params:     []*manager.Param{},
		RouteNames: []string{"pi"},
	}
)

type Response struct {
	Value string `json:"value"`
}

func (r *Response) Text() string {
	return r.Value
}

func (r *Response) Html() string {
	return templates.Render(
		"templates/fragments/tools/single.html",
		templates.C{
			"desc":  "The value of Pi to 60 decimal places is:",
			"value": r.Text(),
		},
	)
}

type Pi struct{}

func New() *Pi {
	return &Pi{}
}

func (p *Pi) Meta() *manager.Meta {
	return meta
}

func (p *Pi) Run(inp *manager.Input) manager.Output {
	return &Response{
		Value: "3.14159265358979323846264338327950288419716939937510582097494",
	}
}
