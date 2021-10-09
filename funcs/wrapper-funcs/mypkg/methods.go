package mypkg

import "fmt"

type MyStruct struct{}

func (ms MyStruct) MyMethod1(s string) string {
	return fmt.Sprint("MyMethod1 stuff", ` - `, s)
}

// func (ms *MyStruct) MyMethod2(s string) string {
// 	return fmt.Sprint("MyMethod2 stuff", ` - `, s)
// }

// Wrapping normal structure method
type CustomMethodType1 func(MyStruct, string) string

func Wrapper2(cm CustomMethodType1) CustomMethodType1 {
	fmt.Println("Wrapper2", ` - `, "START")
	defer fmt.Println("Wrapper2", ` - `, "END")

	return func(ms MyStruct, s string) string {
		fmt.Println("Wrapper2", ` - `, "Added code")
		return cm(ms, s)
	}
}

// // Wrapping reference (pointer) structure method
// type CustomMethodType2 func(*MyStruct, string) string

// // type CustomMethod1 func (*MyStruct) (string) string // This definition causes parsing error
