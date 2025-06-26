package ip

import (
	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/nathan-osman/toolset.sh/templates"
)

var (
	meta = &manager.Meta{
		Name:           "IP Address",
		Desc:           "return the client IP address",
		Params:         []*manager.Param{},
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

func New() *IP {
	return &IP{}
}

func (i *IP) Meta() *manager.Meta {
	return meta
}

func (i *IP) Run(inp *manager.Input) manager.Output {
	return &Response{
		Value: inp.C.ClientIP(),
	}
}
