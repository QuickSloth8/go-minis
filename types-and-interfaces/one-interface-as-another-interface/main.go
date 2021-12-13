/*
This program demonstrates the ability to use the value of one variable, which is of type I1, as type I2,
given the fact that I1 implement I2, and both I1 and I2 are interfaces.
*/

package main

import "fmt"

func main() {
	fmt.Println("Calling as I1 variable:")
	var i1 I1 = S{}
	i1.Print1()
	i1.Print2()

	fmt.Println()

	fmt.Println("Calling as I2 variable:")
	var i2 I2 = i1
	i2.Print1()
}

type I1 interface {
	Print1()
	Print2()
}

type I2 interface {
	Print1()
}

type S struct {
}

func (s S) Print1() {
	fmt.Println("s.Print1")
}

func (s S) Print2() {
	fmt.Println("s.Print2")
}
