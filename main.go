package main

import "fmt"

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}

type InterfaceHere interface {
	Sum() int
}

func main() {
	var a InterfaceHere
	sh := structHere{N1: 1, N2: 2}
	os := otherStruct{A: 2, B: 3}

	a = &sh
	fmt.Println(a.Sum())
	a = os
	fmt.Println(a.Sum())

	var i InterfaceHere = &structHere{N1: 4, N2: 6}
	fmt.Println(i.Sum())
	var o InterfaceHere = otherStruct{A: 4, B: 3}
	fmt.Printf("%v, %T\n", o, o)
	fmt.Println(o.Sum())

	var ai interface{}
	ai = "jelly"
	fmt.Println(ai)
	fmt.Printf("%v, %T\n", ai, ai)
	ai = 42
	fmt.Println(ai)
	fmt.Printf("%v, %T\n", ai, ai)

	ais, ok := ai.(int) // type assertion
	fmt.Println(ais, ok)
	fmt.Printf("%v, %T\n", ais, ais)

	var b interface{} = 3.14

	switch b.(type) {
	case string:
		fmt.Println("b is string")
	case int:
		fmt.Println("b is int")
	case bool:
		fmt.Println("b is bool")
	default:
		fmt.Printf("unknown type %T\n", b)
	}
}
