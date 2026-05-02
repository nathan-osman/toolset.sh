package base64

import (
	"encoding/base64"

	"github.com/nathan-osman/toolset.sh/registry"
	"github.com/nathan-osman/toolset.sh/templates"
	"github.com/nathan-osman/toolset.sh/util"
)

const (
	paramText    = "text"
	paramAction  = "action"
	actionEncode = "encode"
	actionDecode = "decode"
)

var (
	meta = &registry.Meta{
		Category: registry.CategoryProgramming,
		Name:     "Base64",
		Desc:     "encode / decode base64 data",
		Params: []*registry.Param{
			{
				Name:    paramText,
				Desc:    "text to encode / decode",
				Default: "",
			},
			{
				Name:    paramAction,
				Desc:    "action to perform",
				Default: actionEncode,
				Options: []*registry.Option{
					{
						Name:  actionEncode,
						Label: "Encode",
					},
					{
						Name:  actionDecode,
						Label: "Decode",
					},
				},
			},
		},
		RouteName: "base64",
	}
)

type Response struct {
	action string
	Value  string `json:"value"`
}

func (r *Response) Text() string {
	return r.Value
}

func (r *Response) Html() string {
	var desc string
	switch r.action {
	case actionEncode:
		desc = "Your encoded text is:"
	case actionDecode:
		desc = "Your decoded text is:"
	}
	return templates.Render(
		"templates/fragments/tools/single.html",
		templates.C{
			"desc":  desc,
			"value": r.Text(),
			"small": true,
		},
	)
}

type Base64 struct{}

func init() {
	registry.Register(&Base64{})
}

func (b *Base64) Meta() *registry.Meta {
	return meta
}

func (b *Base64) Run(i *registry.Input) registry.Output {
	var (
		action  = util.GetStringParam(i.Params, paramAction, actionEncode)
		vInput  = util.GetStringParam(i.Params, paramText, "")
		vOutput string
	)
	switch action {
	case actionEncode:
		vOutput = base64.StdEncoding.EncodeToString([]byte(vInput))
	case actionDecode:
		v, err := base64.StdEncoding.DecodeString(vInput)
		if err != nil {
			panic(err)
		}
		vOutput = string(v)
	}
	return &Response{
		Value: vOutput,
	}
}
