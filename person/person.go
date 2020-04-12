package person

type Person struct {
	Name     string "json:name"
	Lastname string "json:lastname"
	Age      int    "json:age"
}

func NewPerson(name, lastname string, age int) Person {

	return Person{
		name,
		lastname,
		age,
	}
}
