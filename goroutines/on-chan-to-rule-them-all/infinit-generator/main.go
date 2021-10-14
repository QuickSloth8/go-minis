/*
This program demonstrates an infinite producer,
that outputs values to a channel for the consumer.
The responsibility of closing the channel (stopping
the producer) is on the consumer.
*/

package main

import (
	"log"
)

// GenerateNumbersInBackground Returns a channel that produces numbers incrementally
// starting from 0.
// To stop the underlying goroutine, send a number on the
// returned channel, or close it.
// Note: in order for this to work, the size of the
// channel buffer should be zero.
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
	numbers <- 0
	// Another option, is to close the channel with `close(numbers)`

	/* log.Println(<-numbers) // fatal error: all goroutines are asleep - deadlock! */

	log.Println("Done")
}
