package main

import (
	"fmt"
)

type s struct {
	Key string
}

func main() {
	CompareStruct()
}
func CompareStruct() {
	s1 := s{
		Key: "1",
	}
	s2 := s{
		Key: "1",
	}
	fmt.Printf("%+v == %+v? %v\n", s1, s2, s1 == s2)
	fmt.Printf("%+v == %+v? %v\n", s1, s2, &s1 == &s2)
}
