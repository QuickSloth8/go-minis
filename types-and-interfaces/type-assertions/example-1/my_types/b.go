package my_types

import "fmt"

type B struct {
	Name string
}

func (b *B) F() {
	fmt.Println(b.Name)
}
