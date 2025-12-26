package problems

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Description() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func (p *Person) SetAge() {
	p.Age++
}

func Struct_example() {
	person := Person{
		Name: "John",
		Age:  30,
	}

	fmt.Println(person.Description())
	person.SetAge()
	fmt.Println(person.Description())
	fmt.Println("Age = ", person.Age)
}
