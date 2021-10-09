package mypkg

import "fmt"

func Foo1(s string) string {
	return fmt.Sprint("Foo1 stuff", ` - `, s)
}

type CustomFunctionType func(string) string

func FunctionWrapper(cf CustomFunctionType) CustomFunctionType {
	fmt.Println("FunctionWrapper", ` - `, "START")
	defer fmt.Println("FunctionWrapper", ` - `, "END")

	editedFunc := func(s string) string {
		fmt.Println("FunctionWrapper", ` - `, "Added code")
		return cf(s)
	}
	return editedFunc
}
