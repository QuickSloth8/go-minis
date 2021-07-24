package my_types

import "fmt"

type A struct {
	Name string
}

func (a *A) F() {
	fmt.Println(a.Name)
}
