package main

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	language string
}

func main() {
	//changePriviteValue()
	//unsafeMap()
	transStruct()
}

//利用uintptr和unsafe.Pointer修改成员变量
func changePriviteValue() {
	p := Programmer{"stefno", "go"}
	fmt.Println(p)
	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"
	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"
	fmt.Println(p)
}

func unsafeMap() {
	m := make(map[string]string, 12)
	m["aa"] = "aa"
	m["bb"] = "bb"
	l := **(**int)(unsafe.Pointer(&m))
	fmt.Println(l)
	point := uintptr(unsafe.Pointer(&m))
	count := unsafe.Pointer(point)
	fmt.Println(**(**int)(count))
	flag := unsafe.Pointer(point + unsafe.Sizeof(int(0)))
	fmt.Println(**(**uint8)(flag))
	//B := unsafe.Pointer(point + unsafe.Sizeof(int(0)) + unsafe.Sizeof(uint8(0)) + unsafe.Sizeof(uint8(0)))
	//fmt.Println(**(**uint8)(B))

}

type structA struct {
	Abc string
	Bcd *string
	Cde string
}

type structB struct {
	Abc string
	Bcd *string
}

func transStruct() {
	s := "bcd"
	a := structA{
		Abc: "123",
		Bcd: &s,
	}
	fmt.Println(&a.Abc)
	fmt.Println(a.Bcd)
	b := *(*structB)(unsafe.Pointer(&a))
	fmt.Println(b)
	fmt.Println(&b.Abc)
	fmt.Println(b.Bcd)

}
