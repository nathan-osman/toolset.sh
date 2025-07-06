package useragent

import (
	"github.com/mileusna/useragent"
	"github.com/nathan-osman/toolset.sh/registry"
	"github.com/nathan-osman/toolset.sh/templates"
	"github.com/nathan-osman/toolset.sh/util"
)

var (
	meta = &registry.Meta{
		Category:       registry.CategoryNetwork,
		Name:           "User Agent",
		Desc:           "return information about the user agent",
		Params:         []*registry.Param{},
		RouteName:      "user-agent",
		AlternateNames: []string{"ua"},
	}
	tmplText = util.CompileTemplate(`{{.Value}}

Browser:    {{.Browser}}
Version:    {{.Version}}
OS:         {{or .OS "-"}}
OS version: {{or .OSVersion "-"}}`)
)

type Response struct {
	Value     string `json:"value"`
	Browser   string `json:"browser"`
	Version   string `json:"version"`
	OS        string `json:"os"`
	OSVersion string `json:"os_version"`
}

func (r *Response) Text() string {
	return util.RenderTemplate(tmplText, r)
}

func (r *Response) Html() string {
	return templates.Render(
		"templates/fragments/tools/useragent.html",
		templates.C{
			"desc":     "Your user-agent is:",
			"value":    r.Value,
			"response": r,
		},
	)
}

type UserAgent struct{}

func init() {
	registry.Register(&UserAgent{})
}

func (u *UserAgent) Meta() *registry.Meta {
	return meta
}

func (u *UserAgent) Run(inp *registry.Input) registry.Output {
	v := useragent.Parse(inp.C.GetHeader("User-Agent"))
	return &Response{
		Value:     v.String,
		Browser:   v.Name,
		Version:   v.Version,
		OS:        v.OS,
		OSVersion: v.OSVersion,
	}
}
