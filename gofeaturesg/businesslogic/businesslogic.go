package businesslogic

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

// NewPerson constructs a new person struct with the given name
func NewPerson(name string) *person {
	// You can safely return a pointer to local variable
	// as a local variable will survive the scope of the function.
	p := person{name: name}
	p.age = 42
	return &p
}
