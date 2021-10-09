package main

import "fmt"

func main() {
	///////////////////////////////////////////////////////////
	// Example-1: method with a normal (not a pointer) receiver
	printSectionTitle("Example-1")

	var v1 Interface1 = &Struct1{}
	fmt.Printf("%T, %+v\n", v1, v1)

	var v2 Interface1 = Struct1{}
	fmt.Printf("%T, %+v\n", v2, v2)

	///////////////////////////////////////////////////////////
	// Example-2: method with a pointer receiver
	// note how the second case is not allowed
	printSectionTitle("Example-2")

	var v3 Interface1 = &Struct2{}
	fmt.Printf("%T, %+v\n", v3, v3)

	//var v4 Interface1 = Struct2{}
	// Cannot use 'Struct2{}' (type Struct2) as the type Interface1 Type does not
	// implement 'Interface1' as the 'Print' method has a pointer receiver

	///////////////////////////////////////////////////////////
	// Example-3: storing values in a variable of type *Interface1 (pointer to interface)
	// it includes a lot of examples, where most of them are error cases.
	printSectionTitle("Example-3")

	var v4 *Interface1 = &v1
	fmt.Printf("%T, %+v\n", v4, v4)

	// var v5 *Interface = v1
	// cannot use v1 (type Interface1) as type *Interface1 in assignment:
	//	*Interface1 is pointer to interface, not interface

	// var v6 *Interface1 = &Struct1{}
	// cannot use &Struct1{} (type *Struct1) as type *Interface1 in assignment:
	//	*Interface1 is pointer to interface, not interface

	// var v7 *Interface1 = &(&Struct1{})
	// 1. cannot use &(&Struct1{}) (type **Struct1) as type *Interface1 in assignment:
	//	*Interface1 is pointer to interface, not interface
	// 2. cannot take the address of &Struct1{}

	var v8 Interface1 = Interface1(Struct1{})
	fmt.Printf("%T, %+v\n", v8, v8)

	// var v9 *Interface1 = &Interface1(Struct1{})
	// cannot take the address of Interface1(Struct1{})

}

func printSectionTitle(title string) {
	fmt.Println()
	fmt.Println(title)
	fmt.Println("---------------------")
}

type Interface1 interface {
	Print()
}

type Struct1 struct{}

func (s Struct1) Print() {
	fmt.Println("Struct2", " - ", "Print")
}

type Struct2 struct{}

func (s *Struct2) Print() {
	fmt.Println("*Struct1", " - ", "Print")
}
