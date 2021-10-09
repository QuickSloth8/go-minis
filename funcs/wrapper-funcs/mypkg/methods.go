package mypkg

import "fmt"

type MyStruct struct{}

func (ms MyStruct) MyMethod1(s string) string {
	return fmt.Sprint("MyMethod1 stuff", ` - `, s)
}

func (ms *MyStruct) MyMethod2(s string) string {
	return fmt.Sprint("MyMethod2 stuff", ` - `, s)
}

// CustomMethodType1 defines the type of the method we want to wrap.
// notice that the struct is actually passed as the first parameter
type CustomMethodType1 func(MyStruct, string) string

// MethodWrapper1 does the wrapping of a normal structure method
func MethodWrapper1(cm CustomMethodType1) CustomMethodType1 {
	fmt.Println("MethodWrapper1", ` - `, "START")
	defer fmt.Println("MethodWrapper1", ` - `, "END")

	return func(ms MyStruct, s string) string {
		fmt.Println("MethodWrapper1", ` - `, "Added code")
		return cm(ms, s)
	}
}

// Wrapping reference (pointer) structure method

// CustomMethodType2 defines the type of the method we want to wrap.
// notice that the struct reference is actually passed as the first parameter
type CustomMethodType2 func(*MyStruct, string) string

//MethodWrapper2 does the wrapping of a reference structure method
func MethodWrapper2(cm CustomMethodType2) CustomMethodType2 {
	fmt.Println("MethodWrapper2", ` - `, "START")
	defer fmt.Println("MethodWrapper2", ` - `, "END")

	return func(ms *MyStruct, s string) string {
		fmt.Println("MethodWrapper2", ` - `, `Added code`)
		return cm(ms, s)
	}
}
