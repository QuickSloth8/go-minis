/*
Demonstrates the behaviour on collisions between embedded types.

In case of matching method/variable named, the embedded type should
be explicitly stated.

A note about similar case variations:
You could also test similar cases, where you embed pointers, or mix
pointers with normal embedding.
Also, you can try matching method name between two types, but where one takes a pointer,
and the other a value.

But in all these cases, the compiler still throws an error, because of the ambiguity,
caused by the matching method name.
*/

package main

import "fmt"

func main() {
	a := &Struct3{
		Struct1: Struct1{"A1", 1},
		Struct2: Struct2{"A2", 2},
	}
	fmt.Println(a.B)
	fmt.Println(a.C)
	//fmt.Println(a.A) // compilation error - ambiguous selector a.A
	fmt.Println(a.Struct1.A)
	fmt.Println(a.Struct2.A)

	a.bar1()
	a.bar2()
	//a.foo() // compilation error - ambiguous selector a.foo
	a.Struct1.foo()
	a.Struct2.foo()

}

type Struct1 struct {
	A string
	B int
}

func (s Struct1) foo() {
	fmt.Println("Struct1's foo()")
}

func (s Struct1) bar1() {
	fmt.Println("bar1()")
}

type Struct2 struct {
	A string
	C int
}

func (s Struct2) foo() {
	fmt.Println("Struct2's foo()")
}

func (s Struct2) bar2() {
	fmt.Println("bar2()")
}

type Struct3 struct {
	Struct1
	Struct2
}
