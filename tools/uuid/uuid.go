package uuid

import (
	"github.com/gofrs/uuid/v5"
	"github.com/nathan-osman/toolset.sh/manager"
)

type Response struct {
	Value string `json:"value"`
}

func (r *Response) Text() string {
	return r.Value
}

func (r *Response) Html() string {
	return r.Text()
}

type Uuid struct{}

func New() *Uuid {
	return &Uuid{}
}

func (u *Uuid) Name() string {
	return "UUID v4"
}

func (u *Uuid) Desc() string {
	return "generate a v4 UUID"
}

func (u *Uuid) RouteNames() []string {
	return []string{"uuid", "uuid4"}
}

func (u *Uuid) Run() (manager.Output, error) {
	v, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &Response{
		Value: v.String(),
	}, nil
}
