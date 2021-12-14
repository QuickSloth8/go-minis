/*
This program demonstrates the behaviour of "new" keyword.
*/

package main

import "fmt"

func main() {
	type S struct {
		num int
	}

	a := S{1}
	fmt.Printf("%T, %#v\n", a, a)

	b := &S{2}
	fmt.Printf("%T, %#v\n", b, b)

	c := new(S)
	fmt.Printf("%T, %#v\n", c, c)

	d := new(*S)
	fmt.Printf("%T, %#v\n", d, d)
}
