package people

type Person struct {
	Data PersonData `json:"data"`
}

type PersonData struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Name string `json:"name"`
}
