package uuid

type Uuid struct {
	Type string `json:"type"`
}

func (u *Uuid) Name() string {
	return "UUID v4"
}

func (u *Uuid) Desc() string {
	return "generate a v4 UUID"
}

func (u *Uuid) Commands() []string {
	return []string{"uuid", "uuid4"}
}

func (u *Uuid) Run() (string, string, error) {
	return "", "", nil
}
