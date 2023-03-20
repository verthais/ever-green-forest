package characters

func NewAttribute(name string, bounds []int) Attribute {
	return Attribute{
		Name:  name,
		Value: getRandom(bounds[0], bounds[1]),
	}
}

type Attribute struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
