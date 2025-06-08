package ip

import (
	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/nathan-osman/toolset.sh/util"
)

var (
	t    = util.MustCreateTemplate(`<div class="text-2xl">{{.}}</div>`)
	meta = &manager.Meta{
		Name:       "IP Address",
		Desc:       "return the client IP address",
		Params:     []*manager.Param{},
		RouteNames: []string{"ip"},
	}
)

type Response struct {
	Value string `json:"value"`
}

func (r *Response) Text() string {
	return r.Value
}

func (r *Response) Html() string {
	return util.MustRenderTemplateToString(t, r.Text())
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
