/*
This program demonstrates how to check if a type implements certain methods.
It's possible by type assertion, with the help of an anonymous interface declaration.
*/

package main

import "fmt"

func main() {
	var i1 I1 = S1{}
	a, ok := i1.(interface {
		Print1()
	})
	fmt.Println("ok:", ok)
	fmt.Printf("%T\n\n", a)

	b, ok := i1.(interface {
		Print3()
	})
	fmt.Println("ok:", ok)
	fmt.Printf("%T\n\n", b)
}

type I1 interface {
	Print1()
	Print2()
}

type S1 struct {
}

func (s S1) Print1() {
	fmt.Println("S1.Print1")
}
func (s S1) Print2() {
	fmt.Println("S1.Print2")
}
