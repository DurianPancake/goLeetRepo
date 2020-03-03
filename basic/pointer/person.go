package pointer

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p *Person) SayName() {
	fmt.Println("My name is", p.Name)
	person := Person{
		Name: "张三",
		age:  19,
	}
	*p = person
}
