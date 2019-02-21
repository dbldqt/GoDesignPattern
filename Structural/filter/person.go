package filter

type Person struct {
	name string
	gender string
	maritalStatus string
}

func (per *Person)GetName() string{
	return per.name
}

func (per *Person)GetGender() string{
	return per.gender
}

func (per *Person)GetMaritalStatus() string{
	return per.maritalStatus
}

func NewPerson(name,gender,maritalStatus string) Person{
	return Person{
		name,
		gender,
		maritalStatus,
	}
}


