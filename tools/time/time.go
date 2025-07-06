package time

import (
	"fmt"
	"time"

	"github.com/nathan-osman/toolset.sh/registry"
	"github.com/nathan-osman/toolset.sh/templates"
	"github.com/nathan-osman/toolset.sh/util"
)

const (
	paramFormat = "format"

	formatDefault = "default"
	formatISO8601 = "iso8601"
	formatUnix    = "unix"

	timeFormatDefault = "Monday, January 2, 2006  3:04:05 PM"
	timeFormatISO8601 = "2006-01-02T15:04:05.999Z"
)

var (
	meta = &registry.Meta{
		Category: registry.CategoryDateTime,
		Name:     "Current Time",
		Desc:     "return the current date and time",
		Params: []*registry.Param{
			{
				Name:    paramFormat,
				Desc:    "format for date / time",
				Default: formatDefault,
				Options: []*registry.Option{
					{
						Name:  formatDefault,
						Label: "default",
					},
					{
						Name:  formatISO8601,
						Label: "ISO 8601",
					},
				},
			},
		},
		RouteName:      "time",
		AlternateNames: []string{},
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
			"desc":  "The current UTC date / time is:",
			"value": r.Text(),
		},
	)
}

type Time struct{}

func init() {
	registry.Register(&Time{})
}

func (t *Time) Meta() *registry.Meta {
	return meta
}

func (t *Time) Run(i *registry.Input) registry.Output {
	var (
		n = time.Now()
		f = util.GetStringParam(i.Params, paramFormat, formatDefault)
		s string
	)
	switch f {
	case formatDefault:
		s = n.Format(timeFormatDefault)
	case formatISO8601:
		s = n.Format(timeFormatISO8601)
	case formatUnix:
		s = fmt.Sprintf("%d", n.Unix())
	default:
		panic("invalid value for parameter 'format'")
	}
	return &Response{
		Value: s,
	}
}
