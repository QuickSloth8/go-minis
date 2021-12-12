package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//TODO: describe each Ex - for each, describe when and if the issue will be detected, what's the issue, and how to fix it
}

// Ex1 using a loop variable
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

// Ex2 a valid solution, where a new j is created each loop
func Ex2() {
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

// Ex3 an issue when reusing a variable created outside the loop
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
