/*
The program demonstrates how to wrap functions & methods

Note: wrapper code executed at wrap time, once, instead of every
call to the modified function. While the modified function code
runs every time.
*/
package main

import (
	"fmt"
	"gominis/mypkg"
)

func main() {
	// var _ mypkg.CustomMethodType1 = (mypkg.MyStruct).MyMethod1
	// var _ CustomMethodType1 = (&MyStruct{}).MyMethod2

	s := "JustAString"

	// // Example 1: wrapping a normal function
	// fmt.Println(mypkg.Foo1(s))
	// displaySpacer()

	// wrapped1 := mypkg.Wrapper1(mypkg.Foo1)
	// fmt.Println("Here is the modified function ...")
	// fmt.Println(wrapped1(s))
	// displaySpacer()

	// Example 2: wrapping normal method
	ms := mypkg.MyStruct{}
	fmt.Println(ms.MyMethod1(s))
	displaySpacer()

	wrapped2 := mypkg.Wrapper2((mypkg.MyStruct).MyMethod1)
	fmt.Println(wrapped2(ms, s))
	displaySpacer()
}

func displaySpacer() {
	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println()
}
