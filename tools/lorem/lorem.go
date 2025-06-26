package lorem

import (
	"errors"
	"strings"

	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/nathan-osman/toolset.sh/templates"
	"github.com/nathan-osman/toolset.sh/util"
)

const (
	paramNum = "num"
)

var (
	meta = &manager.Meta{
		Name: "Lorem Ipsum",
		Desc: "generate random text",
		Params: []*manager.Param{
			{
				Name:    paramNum,
				Desc:    "number of paragraphs to generate",
				Default: "3",
			},
		},
		RouteName:      "lorem-ipsum",
		AlternateNames: []string{"lorem"},
	}
	paragraphs = []string{
		"Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Ut purus elit, vestibulum ut, placerat ac, adipiscing vitae, felis. Curabitur dictum gravida mauris. Nam arcu libero, nonummy eget, consectetuer id, vulputate a, magna. Donec vehicula augue eu neque. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Mauris ut leo. Cras viverra metus rhoncus sem. Nulla et lectus vestibulum urna fringilla ultrices. Phasellus eu tellus sit amet tortor gravida placerat. Integer sapien est, iaculis in, pretium quis, viverra ac, nunc. Praesent eget sem vel leo ultrices bibendum. Aenean faucibus. Morbi dolor nulla, malesuada eu, pulvinar at, mollis ac, nulla. Curabitur auctor semper nulla. Donec varius orci eget risus. Duis nibh mi, congue eu, accumsan eleifend, sagittis quis, diam. Duis eget orci sit amet orci dignissim rutrum.",
		"Nam dui ligula, fringilla a, euismod sodales, sollicitudin vel, wisi. Morbi auctor lorem non justo. Nam lacus libero, pretium at, lobortis vitae, ultricies et, tellus. Donec aliquet, tortor sed accumsan bibendum, erat ligula aliquet magna, vitae ornare odio metus a mi. Morbi ac orci et nisl hendrerit mollis. Suspendisse ut massa. Cras nec ante. Pellentesque a nulla. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aliquam tincidunt urna. Nulla ullamcorper vestibulum turpis. Pellentesque cursus luctus mauris.",
		"Nulla malesuada porttitor diam. Donec felis erat, congue non, volutpat at, tincidunt tristique, libero. Vivamus viverra fermentum felis. Donec nonummy pellentesque ante. Phasellus adipiscing semper elit. Proin fermentum massa ac quam. Sed diam turpis, molestie vitae, placerat a, molestie nec, leo. Maecenas lacinia. Nam ipsum ligula, eleifend at, accumsan nec, suscipit a, ipsum. Morbi blandit ligula feugiat magna. Nunc eleifend consequat lorem. Sed lacinia nulla vitae enim. Pellentesque tincidunt purus vel magna. Integer non enim. Praesent euismod nunc eu purus. Donec bibendum quam in tellus. Nullam cursus pulvinar lectus. Donec et mi. Nam vulputate metus eu enim. Vestibulum pellentesque felis eu massa.",
	}
)

type Response struct {
	Paragraphs []string `json:"paragraphs"`
}

func (r *Response) Text() string {
	return strings.Join(r.Paragraphs, "\n\n")
}

func (r *Response) Html() string {
	return templates.Render(
		"templates/fragments/tools/single.html",
		templates.C{
			"desc":  "Your generated text is:",
			"value": r.Text(),
			"small": true,
		},
	)
}

type Lorem struct{}

func New() *Lorem {
	return &Lorem{}
}

func (l *Lorem) Meta() *manager.Meta {
	return meta
}

func (l *Lorem) Run(i *manager.Input) manager.Output {
	var (
		n = util.GetIntParam(i.Params, paramNum, 3)
		v = paragraphs
	)
	if n < 0 || n > 100 {
		panic(errors.New("num must be between 0 and 100"))
	}
	for i := 0; i < n; i += 3 {
		v = append(v, paragraphs...)
	}
	return &Response{
		Paragraphs: v[:n],
	}
}
