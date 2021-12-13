/*
The program includes simple data-race cases, which might happen in loops.
These data-races might cause unexpected read values on certain variable.

You can find additional cases at https://go.dev/doc/articles/race_detector.

Tip: run with `go run -race ...` to detect data-races (assuming a call to the function which causes the data-race).
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Ex1:")
	Ex1()
	fmt.Println()

	fmt.Println("Ex2:")
	Ex2()
	fmt.Println()

	fmt.Println("Ex3:")
	Ex3()
	fmt.Println()

	fmt.Println("Ex4:")
	Ex4()
	fmt.Println()
}

// Ex1 using a loop variable inside an inner sub-scope
func Ex1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			fmt.Println(i)
		}()
	}

	wg.Wait()
}

// Ex2 solving issue from Ex1, by passing loop-variable to goroutine, which makes a copy of it.
func Ex2() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}

// Ex3 has a data-race, due to reusing a variable, which is created in outer scope, in different goroutines.
func Ex3() {
	wg := sync.WaitGroup{}
	var j int
	for i := 0; i < 3; i++ {
		j = i
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			fmt.Println(j)
		}()
	}

	wg.Wait()
}

// Ex4 solving issue from Ex3, by reinitializing a new j in each iteration.
func Ex4() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			fmt.Println(j)
		}()
	}

	wg.Wait()
}
