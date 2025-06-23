package uuid

import (
	"github.com/gofrs/uuid/v5"
	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/nathan-osman/toolset.sh/templates"
	"github.com/nathan-osman/toolset.sh/util"
)

const (
	paramType = "type"

	typeUuid4 = "uuid4"
	typeUuid7 = "uuid7"
)

var (
	meta = &manager.Meta{
		Name: "Generate UUID",
		Desc: "generate a UUID (universally unique identifier)",
		Params: []*manager.Param{
			{
				Name:    paramType,
				Desc:    "type of UUID",
				Default: typeUuid4,
				Options: []*manager.Option{
					{
						Name:  typeUuid4,
						Label: "UUID version 4",
					},
					{
						Name:  typeUuid7,
						Label: "UUID version 7",
					},
				},
			},
		},
		RouteNames: []string{"uuid", "uuid4"},
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
			"desc":  "Your randomly generated UUID is:",
			"value": r.Text(),
		},
	)
}

type Uuid struct{}

func New() *Uuid {
	return &Uuid{}
}

func (u *Uuid) Meta() *manager.Meta {
	return meta
}

func (u *Uuid) Run(i *manager.Input) manager.Output {
	var (
		v   uuid.UUID
		err error
		t   = util.GetStringParam(i.Params, paramType, typeUuid4)
	)
	switch t {
	case typeUuid4:
		v, err = uuid.NewV4()
	case typeUuid7:
		v, err = uuid.NewV7()
	default:
		panic("invalid value for parameter 'type'")
	}
	if err != nil {
		panic(err)
	}
	return &Response{
		Value: v.String(),
	}
}
