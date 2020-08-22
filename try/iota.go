package main

import "fmt"

const (
	TypeA = 1 << iota
	TypeB
	TypeC
	TypeD
)
const (
	TypeAA = iota
	TypeBB
	TypeCC
	TypeDD
)

func main() {
	fmt.Printf("%+v\n", TypeA)
	fmt.Printf("%+v\n", TypeB)
	fmt.Printf("%+v\n", TypeC)
	fmt.Printf("%+v\n", TypeD)

	fmt.Printf("%+v\n", TypeAA)
	fmt.Printf("%+v\n", TypeBB)
	fmt.Printf("%+v\n", TypeCC)
	fmt.Printf("%+v\n", TypeDD)
}
