package fact

// Fact represents the fact of social
// relation between two persons
type Fact struct {
	A             Person        `json:"person_a"`
	B             Person        `json:"person_b"`
	Communication Communication `json:"communication"`
}

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func (p Person) IsValid() bool {
	return p.Name != "" && p.Surname != "" && p.Age != 0
}

type Communication struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (c Communication) DescriptionNoNEmpty() bool {
	return c.Description != ""
}
