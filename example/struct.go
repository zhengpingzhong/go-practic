package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "alex", age: 30})
	fmt.Println(person{age: 34})
	fmt.Println(&person{name: "anna", age: 40})
	fmt.Println(newPerson("john"))
	s := person{name: "sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 52
	fmt.Println(sp.age)
}
