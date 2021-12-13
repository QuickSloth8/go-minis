/*
Demonstrates the special case of a nil pointer to struct
Some limited method calls are allowed, but others not, and they cause a panic.

Notice the two Examples: Example1 & Example2 ...

In Example1, the pointer is kept as nil, and this causes a panic in most cases.
Only case allowed is calling a method that takes a pointer to struct, and doesn't
access any member variables of the struct.

In Example2, the pointer is initialized the 'new' keyword.
In this case, as such in the case of initializing with &A{}, all calls are allowed,
since the pointer is no longer nil, and it's fields initialized with defaults.
*/

package main

import (
	"fmt"
)

func main() {
	Example1()
	Example2()
}

// Example1 - a nil pointer to struct
func Example1() {
	fmt.Println("## Example1 ##")
	var pointerA *A
	fmt.Printf("type: %T, value: %+v\n", pointerA, pointerA)

	fmt.Println("(*) calling pointer method, that doesn't access member var")
	pointerA.PrintSomething()
	fmt.Println("DONE")

	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		the panic occurs after calling the method, and only when trying the access pointerA member var
	*/
	//fmt.Println("(*) calling pointer method, that DOES access member var")
	//pointerA.PrintMemberValue()

	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		the panic occurs before the call, when trying to take the value of pointerA
	*/
	//fmt.Println("(*) calling value method, that doesn't access member var")
	//pointerA.PrintSomething2()

	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		the panic occurs before the call, when trying to take the value of pointerA
	*/
	//fmt.Println("(*) calling value method, that DOES access member var")
	//pointerA.PrintMemberValue2()
}

// Example2 - initializing with keyword new
func Example2() {
	fmt.Println("## Example2 ##")
	pointerA := new(A)
	fmt.Printf("type: %T, value: %+v\n", pointerA, pointerA)

	fmt.Println("(*) calling pointer method, that doesn't access member var")
	pointerA.PrintSomething()
	fmt.Println("DONE")

	fmt.Println("(*) calling pointer method, that DOES access member var")
	pointerA.PrintMemberValue()
	fmt.Println("DONE")

	fmt.Println("(*) calling value method, that doesn't access member var")
	pointerA.PrintSomething2()
	fmt.Println("DONE")

	fmt.Println("(*) calling value method, that DOES access member var")
	pointerA.PrintMemberValue2()
	fmt.Println("DONE")
}

type A struct {
	S string
}

// Pointer Methods

func (a *A) PrintSomething() {
	fmt.Println("PrintSomething")
}

func (a *A) PrintMemberValue() {
	fmt.Println(a.S)
}

// Value Methods

func (a A) PrintSomething2() {
	fmt.Println("PrintSomething")
}

func (a A) PrintMemberValue2() {
	fmt.Println(a.S)
}
