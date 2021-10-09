/*
The program demonstrates how to wrap functions & methods.
A wrapper is a function/method that take a function/method type,
does something, and returns a new function/method of the same type.

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
	// validate same type
	var _ mypkg.CustomMethodType1 = mypkg.MyStruct.MyMethod1
	var _ mypkg.CustomMethodType2 = (*mypkg.MyStruct).MyMethod2

	s := "JustAString"

	// Example 1: wrapping a normal function
	displaySectionTitle("Example of function wrapping:")

	displaySubSectionTitle("Original function output:")
	fmt.Println(mypkg.Foo1(s))
	displaySpacer()

	displaySubSectionTitle("Wrapping output:")
	wrapped1 := mypkg.FunctionWrapper(mypkg.Foo1)
	displaySpacer()

	fmt.Println("Modified function output:")
	fmt.Println(wrapped1(s))
	displaySpacer()

	// Example 2: wrapping normal method (doesn't require struct reference/pointer)
	displaySectionTitle("Example of normal method wrapping:")

	ms1 := mypkg.MyStruct{}
	displaySubSectionTitle("Original method output:")
	fmt.Println(ms1.MyMethod1(s))
	displaySpacer()

	displaySubSectionTitle("Wrapping output:")
	wrappedMethod1 := mypkg.MethodWrapper1(mypkg.MyStruct.MyMethod1)
	displaySpacer()

	displaySubSectionTitle("Modified method output:")
	fmt.Println(wrappedMethod1(ms1, s))
	displaySpacer()

	// Example 3: wrapping reference method
	displaySectionTitle("Example of reference method wrapping:")

	ms2 := &mypkg.MyStruct{}
	displaySubSectionTitle("Original method output:")
	fmt.Println(ms2.MyMethod2(s))
	displaySpacer()

	displaySubSectionTitle("Wrapping output:")
	wrappedMethod2 := mypkg.MethodWrapper2((*mypkg.MyStruct).MyMethod2)
	displaySpacer()

	displaySubSectionTitle("Modified method output:")
	fmt.Println(wrappedMethod2(ms2, s))
	displaySpacer()
}

func displaySpacer() {
	fmt.Println("--------------------------")
}

func displaySectionTitle(title string) {
	fmt.Println()
	fmt.Println(`###########################`)
	fmt.Println(`# ` + title)
	fmt.Println(`###########################`)
	fmt.Println()
}

func displaySubSectionTitle(title string) {
	fmt.Println(`(*) ` + title)
}
