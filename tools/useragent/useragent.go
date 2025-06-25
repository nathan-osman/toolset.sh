package useragent

import (
	"github.com/mileusna/useragent"
	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/nathan-osman/toolset.sh/templates"
	"github.com/nathan-osman/toolset.sh/util"
)

var (
	meta = &manager.Meta{
		Name:       "User Agent",
		Desc:       "return information about the user agent",
		Params:     []*manager.Param{},
		RouteNames: []string{"user-agent", "useragent", "ua"},
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

func New() *UserAgent {
	return &UserAgent{}
}

func (u *UserAgent) Meta() *manager.Meta {
	return meta
}

func (u *UserAgent) Run(inp *manager.Input) manager.Output {
	v := useragent.Parse(inp.C.GetHeader("User-Agent"))
	return &Response{
		Value:     v.String,
		Browser:   v.Name,
		Version:   v.Version,
		OS:        v.OS,
		OSVersion: v.OSVersion,
	}
}
