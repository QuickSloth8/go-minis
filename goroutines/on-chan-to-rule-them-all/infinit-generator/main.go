package main

import (
	"log"
)

// Returns a channel that produces numbers incrementally
// starting from 0
// To stop the underlying goroutine, send a number on the
// returned channel, or close it
func GenerateNumbersInBackground() chan int {
	c := make(chan int)

	go func() {
		n := 0
		for {
			select {
			case c <- n:
				n++
			case <-c:
				log.Println("terminating goroutine ...")
			}
		}
	}()

	return c
}

func main() {

	numbers := GenerateNumbersInBackground()

	for i := 0; i < 10; i++ {
		log.Println(<-numbers)
	}

	log.Println("closing channel")
	close(numbers)
	/* log.Println(<-numbers) // fatal error: all goroutines are asleep - deadlock! */
	log.Println("Done")

}
