package main

import (
	"app/limit"
	"fmt"
	"sync"
)

func main() {
	limiter := limit.NewLimiter(4, func(message string) {
		fmt.Println(message)
	})

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			limiter.ExecBegin()
			fmt.Printf("doing work %d\n", i)
			if err := limiter.ExecEnd(); err != nil {
				panic(err)
			}
		}(i)
	}

	wg.Wait()
}
