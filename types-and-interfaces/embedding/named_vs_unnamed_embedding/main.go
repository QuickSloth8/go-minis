/*
Demonstrated the difference between named & unnamed embeddings,
especially when calling methods/referencing member variables.
*/

package main

import (
	"app/my_structs"
	"fmt"
)

func main() {
	s := my_structs.NewStruct4()

	// In both calls, Struct1 is linked
	fmt.Println(s.A)
	s.Foo()

	// To access named embeddings, the respective should be used
	fmt.Println(s.S2.A)
	s.S2.Foo()
}
