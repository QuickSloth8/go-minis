/*
Demonstrates how a structure can implement multiple interfaces, and
that it can be verified based solely on its methods.
*/

package main

import (
	"fmt"
	aPkg "gominis/a"
	bPkg "gominis/b"
	"io"
)

type notImplementedI interface {
	Foo()
}

type myWriter struct {
}

func (w myWriter) Write(p []byte) (n int, err error) {
	fmt.Println("Should write", p)
	n = 0
	err = nil
	return n, err
}

func main() {
	var _ aPkg.Writer = (*myWriter)(nil) // Verify that *myWriter implements a.Writer
	fmt.Println("myWriter implements a.Writer")
	var _ bPkg.Writer = (*myWriter)(nil) // Verify that *myWriter implements b.Writer
	fmt.Println("myWriter implements b.Writer")
	var _ io.Writer = (*myWriter)(nil) // verify that *myWriter implements io.Writer
	fmt.Println("myWriter implements io.Writer")
	// var _ notImplementedI = (*myWriter)(nil) // ERROR - cannot use (*myWriter)(nil) (type *myWriter) as type notImplementedI in assignment
}
