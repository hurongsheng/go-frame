package main

import "fmt"

func main() {
	s := student{person{Id: 1, Name: "p", Age: 333}, &child{Id: 2, Name: "c"}, 3, "s"}
	s.do()
	s.person.do()
	s.child.do()
	s.done()
	fmt.Println("name", s.Name)
	fmt.Println("age", s.Age)
	fmt.Println("id", s.Id)
}

type person struct {
	Id   int
	Name string
	Age  int
}

func (p *person) do() {
	fmt.Println("person ", p.Name)
}

func (p *person) done() {
	fmt.Println("person done ", p.Name)
}

type child struct {
	Id   int
	Name string
}

func (c *child) do() {
	fmt.Println("child ", c.Name)
}

type student struct {
	person
	*child
	Id   int
	Name string
}

func (s *student) do() {
	fmt.Println("student do ", s.Name)

}
