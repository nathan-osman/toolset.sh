package rand

import (
	"fmt"
	"math/rand/v2"

	"github.com/nathan-osman/toolset.sh/manager"
)

type Response struct {
	Value float64 `json:"value"`
}

func (r *Response) Text() string {
	return fmt.Sprintf("%f", r.Value)
}

func (r *Response) Html() string {
	return r.Text()
}

type Rand struct {
	rand *rand.Rand
}

func New() *Rand {
	return &Rand{
		rand: rand.New(
			rand.NewPCG(1, 2),
		),
	}
}

func (r *Rand) Name() string {
	return "Random Number"
}

func (r *Rand) Desc() string {
	return "generate a random number"
}

func (r *Rand) RouteNames() []string {
	return []string{"rand", "random"}
}

func (r *Rand) Run() (manager.Output, error) {
	return &Response{
		Value: r.rand.Float64(),
	}, nil
}
