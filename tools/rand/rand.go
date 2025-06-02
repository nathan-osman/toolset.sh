package rand

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/nathan-osman/toolset.sh/util"
)

const (
	paramMin = "min"
	paramMax = "max"
)

var (
	t    = util.MustCreateTemplate(`<div class="text-2xl">{{.}}</div>`)
	meta = &manager.Meta{
		Name: "Random number",
		Desc: "generate a random number",
		Params: []*manager.Param{
			{
				Name:    paramMin,
				Desc:    "minimum value, as a floating point number",
				Default: "0",
			},
			{
				Name:    paramMax,
				Desc:    "maximum value, as a floating point number",
				Default: "1",
			},
		},
		RouteNames: []string{"rand", "random"},
	}
)

type Response struct {
	Value float64 `json:"value"`
}

func (r *Response) Text() string {
	return fmt.Sprintf("%f", r.Value)
}

func (r *Response) Html() string {
	return util.MustRenderTemplateToString(t, r.Text())
}

type Rand struct{}

func New() *Rand {
	return &Rand{}
}

func (r *Rand) Meta() *manager.Meta {
	return meta
}

func (r *Rand) Run(i *manager.Input) manager.Output {
	var (
		b   = make([]byte, 8)
		min = util.GetFloatParam(i.Params, paramMin, 0.0)
		max = util.GetFloatParam(i.Params, paramMax, 1.0)
	)
	rand.Read(b)
	var (
		randInt   = binary.BigEndian.Uint64(b)
		randFloat = float64(randInt) / (1 << 64)
	)
	return &Response{
		Value: min + randFloat*(max-min),
	}
}
