/*
UNFINISHED !!
*/
package main

import (
	"fmt"
	"gominis/types/type-assertions/example-1/my_types"
	"reflect"
)

type I1 interface {
	F()
}

func foo(i []interface{}) {
	for _, e := range i {
		// _, ok := e.(I1)
		// fmt.Println(ok)
		// fmt.Printf("%T, %v\n", e, e)
		fmt.Println(reflect.TypeOf(e))
	}
}

func main() {
	a := &my_types.A{Name: "AAA"}
	b := &my_types.B{Name: "BBB"}

	s := append(make([]interface{}, 0), a, b)

	foo(s)
	// fmt.Println(s)
}
