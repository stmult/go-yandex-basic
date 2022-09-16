package main

import "fmt"

func main() {
	a := new(Android)
	a.Talk()
	p1 := Person{"Ivan"}
	p1.Talk()
	a1 := Android{
		Person: Person{},
		Model:  "XZ",
	}
	a1.Talk()
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}
