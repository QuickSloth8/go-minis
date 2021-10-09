package mypkg

import "fmt"

func Foo1(s string) string {
	return fmt.Sprint("Foo1 stuff", ` - `, s)
}

type CustomFuncType1 func(string) string

func Wrapper1(cf CustomFuncType1) CustomFuncType1 {
	fmt.Println("Wrapper1", ` - `, "START")
	defer fmt.Println("Wrapper1", ` - `, "END")

	editedFunc := func(s string) string {
		fmt.Println("Wrapper1", ` - `, "Added code")
		return cf(s)
	}
	return editedFunc
}
