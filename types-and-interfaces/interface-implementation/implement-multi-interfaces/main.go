/*
Programs demonstrates how a structure can implement multiple interfaces,
based solely on its methods
*/

package main

import (
	"fmt"
	aPkg "gominis/types-and-interfaces/interface-implementation/implement-multi-interfaces/a"
	bPkg "gominis/types-and-interfaces/interface-implementation/implement-multi-interfaces/b"
	"io"
)

type myWriter struct {
}

func (w myWriter) Write(p []byte) (n int, err error) {
	fmt.Println("Should write", p)
	n = 0
	err = nil
	return n, err
}

func main() {
	var _ aPkg.Writer = (*myWriter)(nil) // Verify that *myWriter implements aPkg.Writer
	var _ bPkg.Writer = (*myWriter)(nil) // Verify that *myWriter implements bPkg.Writer
	var _ io.Writer = (*myWriter)(nil)   // verify that *myWriter implements io.Writer
}
