package my_structs

import "fmt"

type Struct1 struct {
	A string
}

func (s Struct1) Foo() {
	fmt.Println("Struct1 - Foo()")
}

type Struct2 struct {
	A string
}

func (s Struct2) Foo() {
	fmt.Println("Struct2 - Foo()")
}

type Struct3 struct {
	A string
}

func (s Struct3) Foo() {
	fmt.Println("Struct3 - Foo()")
}

type Struct4 struct {
	Struct1
	S2 Struct2
	s3 Struct3
}

func NewStruct4() *Struct4 {
	return &Struct4{
		Struct1: Struct1{"A1"},
		S2:      Struct2{"A2"},
		s3:      Struct3{"A3"},
	}
}
