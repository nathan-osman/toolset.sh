package ip

import (
	"github.com/nathan-osman/toolset.sh/registry"
	"github.com/nathan-osman/toolset.sh/templates"
)

var (
	meta = &registry.Meta{
		Category:       registry.CategoryNetwork,
		Name:           "IP Address",
		Desc:           "return the client IP address",
		Params:         []*registry.Param{},
		RouteName:      "ip-address",
		AlternateNames: []string{"ip"},
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
		"templates/fragments/tools/ip.html",
		templates.C{
			"desc":  "Your current IP address is:",
			"value": r.Text(),
		},
	)
}

type IP struct{}

func init() {
	registry.Register(&IP{})
}

func (i *IP) Meta() *registry.Meta {
	return meta
}

func (i *IP) Run(inp *registry.Input) registry.Output {
	return &Response{
		Value: inp.C.ClientIP(),
	}
}
