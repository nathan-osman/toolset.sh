package rand

type Rand struct {
	Min float64
	Max float64
}

func (r *Rand) Name() string {
	return "Random Number"
}

func (r *Rand) Desc() string {
	return "generate a random number"
}

func (r *Rand) Commands() []string {
	return []string{"rand", "random"}
}

func (r *Rand) Run() (string, string, error) {
	return "", "", nil
}
