/*
Demonstrates how a structure can implement multiple interfaces, judging solely on its methods.
Also, an explicit way to verify the implementation of certain interface.
*/

package main

import (
	"fmt"
	aPkg "gominis/a"
	bPkg "gominis/b"
	"io"
)

type strInterface1 interface {
	GetString() string
}

type myWriter struct {
}

type strInterface2 interface {
	GetString() (string, error)
}

func (w myWriter) Write(p []byte) (n int, err error) {
	fmt.Println("Should write", p)
	n = 0
	err = nil
	return n, err
}

func (w myWriter) GetString() string {
	return fmt.Sprintln("myWriter.Print")
}

func main() {
	// note: since we implemented methods on myWriter, instead of *myWriter, *myWriter reference would be implicitly
	//	converted to myWriter.

	var _ aPkg.Writer = (*myWriter)(nil) // verifies that *myWriter implements a.Writer
	fmt.Println("myWriter implements a.Writer")

	var _ bPkg.Writer = (*myWriter)(nil) // verifies that *myWriter implements b.Writer
	fmt.Println("myWriter implements b.Writer")

	var _ io.Writer = (*myWriter)(nil) // verifies that *myWriter implements io.Writer
	fmt.Println("myWriter implements io.Writer")

	var _ strInterface1 = (*myWriter)(nil) // verifies that *myWriter implements strInterface1
	fmt.Println("myWriter implements strInterface1")

	//var _ strInterface2 = (*myWriter)(nil) // proves that *myWriter does not implement strInterface2
	fmt.Println("myWriter does not implement strInterface2")
}
