/*
Inspired by: https://travix.io/type-embedding-in-go-ba40dd4264df
Demonstrates embedding an interfaced in structures.

Notice how Football struct hides the embedded Ball struct, by
embedding Bouncer interface.
*/

package main

import "fmt"

func main() {
	fb := Football{&Ball{Radius: 5, Material: "leather"}}
	fb.Bounce()
}

type Bouncer interface {
	Bounce()
}

type Ball struct {
	Radius   int
	Material string
}

func (b *Ball) Bounce() {
	fmt.Printf("Bouncing ball %+v\n", b)
}

type Football struct {
	Bouncer
}
